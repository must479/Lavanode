package lavasession

import (
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/lavanet/lava/utils"
)

type ProviderSessionManager struct {
	sessionsWithAllConsumers                map[uint64]map[string]*ProviderSessionsWithConsumer // first key is epochs, second key is a consumer address
	dataReliabilitySessionsWithAllConsumers map[uint64]map[string]*ProviderSessionsWithConsumer // separate handling of data reliability so later on we can use it outside of pairing, first key is epochs, second key is a consumer address
	subscriptionSessionsWithAllConsumers    map[uint64]map[string]map[string]*RPCSubscription   // first key is an epoch, second key is a consumer address, third key is subscriptionId
	lock                                    sync.RWMutex
	blockedEpochHeight                      uint64 // requests from this epoch are blocked
	rpcProviderEndpoint                     *RPCProviderEndpoint
	blockDistanceForEpochValidity           uint64 // sessionsWithAllConsumers with epochs older than ((latest epoch) - numberOfBlocksKeptInMemory) are deleted.
}

// reads cs.BlockedEpoch atomically
func (psm *ProviderSessionManager) atomicWriteBlockedEpoch(epoch uint64) {
	atomic.StoreUint64(&psm.blockedEpochHeight, epoch)
}

// reads cs.BlockedEpoch atomically
func (psm *ProviderSessionManager) atomicReadBlockedEpoch() (epoch uint64) {
	return atomic.LoadUint64(&psm.blockedEpochHeight)
}

func (psm *ProviderSessionManager) GetBlockedEpochHeight() uint64 {
	return psm.atomicReadBlockedEpoch()
}

func (psm *ProviderSessionManager) IsValidEpoch(epoch uint64) (valid bool) {
	return epoch > psm.atomicReadBlockedEpoch()
}

// Check if consumer exists and is not blocked, if all is valid return the ProviderSessionsWithConsumer pointer
func (psm *ProviderSessionManager) IsActiveConsumer(epoch uint64, address string) (providerSessionWithConsumer *ProviderSessionsWithConsumer, err error) {
	providerSessionWithConsumer, err = psm.getActiveConsumer(epoch, address)
	if err != nil {
		return nil, err
	}
	return providerSessionWithConsumer, nil // no error
}

func (psm *ProviderSessionManager) getSingleSessionFromProviderSessionWithConsumer(providerSessionWithConsumer *ProviderSessionsWithConsumer, sessionId uint64, epoch uint64, relayNumber uint64) (*SingleProviderSession, error) {
	if providerSessionWithConsumer.atomicReadConsumerBlocked() != notBlockListedConsumer {
		return nil, utils.LavaFormatError("This consumer address is blocked.", nil, &map[string]string{"RequestedEpoch": strconv.FormatUint(epoch, 10), "consumer": providerSessionWithConsumer.consumerAddr})
	}
	// before getting any sessions.
	singleProviderSession, err := psm.getSessionFromAnActiveConsumer(providerSessionWithConsumer, sessionId, epoch) // after getting session verify relayNum etc..
	if err != nil {
		return nil, utils.LavaFormatError("getSessionFromAnActiveConsumer Failure", err, &map[string]string{"RequestedEpoch": strconv.FormatUint(epoch, 10), "sessionId": strconv.FormatUint(sessionId, 10)})
	}
	if singleProviderSession.RelayNum+1 < relayNumber { // validate relay number here, but add only in PrepareSessionForUsage
		return nil, utils.LavaFormatError("singleProviderSession.RelayNum mismatch, session out of sync", SessionOutOfSyncError, &map[string]string{"singleProviderSession.RelayNum": strconv.FormatUint(singleProviderSession.RelayNum+1, 10), "request.relayNumber": strconv.FormatUint(relayNumber, 10)})
	}
	// singleProviderSession is locked at this point.
	return singleProviderSession, err
}

func (psm *ProviderSessionManager) getOrCreateDataReliabilitySessionWithConsumer(address string, epoch uint64, sessionId uint64) (providerSessionWithConsumer *ProviderSessionsWithConsumer, err error) {
	if mapOfDataReliabilitySessionsWithConsumer, consumerFoundInEpoch := psm.dataReliabilitySessionsWithAllConsumers[epoch]; consumerFoundInEpoch {
		if providerSessionWithConsumer, consumerAddressFound := mapOfDataReliabilitySessionsWithConsumer[address]; consumerAddressFound {
			if providerSessionWithConsumer.atomicReadConsumerBlocked() == blockListedConsumer { // we atomic read block listed so we dont need to lock the provider. (double lock is always a bad idea.)
				// consumer is blocked.
				utils.LavaFormatWarning("getActiveConsumer", ConsumerIsBlockListed, &map[string]string{"RequestedEpoch": strconv.FormatUint(epoch, 10), "ConsumerAddress": address})
				return nil, ConsumerIsBlockListed
			}
			return providerSessionWithConsumer, nil // no error
		}
	} else {
		// If Epoch is missing from map, create a new instance
		psm.dataReliabilitySessionsWithAllConsumers[epoch] = make(map[string]*ProviderSessionsWithConsumer)
	}

	// If we got here, we need to create a new instance for this consumer address.
	providerSessionWithConsumer = NewProviderSessionsWithConsumer(address, nil, isDataReliabilityPSWC)
	psm.dataReliabilitySessionsWithAllConsumers[epoch][address] = providerSessionWithConsumer
	return providerSessionWithConsumer, nil
}

// GetDataReliabilitySession fetches a data reliability session
func (psm *ProviderSessionManager) GetDataReliabilitySession(address string, epoch uint64, sessionId uint64, relayNumber uint64) (*SingleProviderSession, error) {
	// validate Epoch
	if !psm.IsValidEpoch(epoch) { // fast checking to see if epoch is even relevant
		utils.LavaFormatError("GetSession", InvalidEpochError, &map[string]string{"RequestedEpoch": strconv.FormatUint(epoch, 10)})
		return nil, InvalidEpochError
	}

	// validate sessionId
	if sessionId > DataReliabilitySessionId {
		return nil, utils.LavaFormatError("request's sessionId is larger than the data reliability allowed session ID", nil, &map[string]string{"sessionId": strconv.FormatUint(sessionId, 10), "DataReliabilitySessionId": strconv.Itoa(DataReliabilitySessionId)})
	}

	// validate RelayNumber
	if relayNumber > DataReliabilityRelayNumber {
		return nil, utils.LavaFormatError("request's relayNumber is larger than the DataReliabilityRelayNumber allowed in Data Reliability", nil, &map[string]string{"relayNumber": strconv.FormatUint(relayNumber, 10), "DataReliabilityRelayNumber": strconv.Itoa(DataReliabilityRelayNumber)})
	}

	// validate active consumer.
	providerSessionWithConsumer, err := psm.getOrCreateDataReliabilitySessionWithConsumer(address, epoch, sessionId)
	if err != nil {
		return nil, utils.LavaFormatError("getOrCreateDataReliabilitySessionWithConsumer Failed", err, &map[string]string{"relayNumber": strconv.FormatUint(relayNumber, 10), "DataReliabilityRelayNumber": strconv.Itoa(DataReliabilityRelayNumber)})
	}

	// singleProviderSession is locked after this method is called unless we got an error
	singleProviderSession, err := providerSessionWithConsumer.getDataReliabilitySingleSession(sessionId, epoch)
	if err != nil {
		return nil, err
	}

	// validate relay number in the session stored
	if singleProviderSession.RelayNum+1 > DataReliabilityRelayNumber { // validate relay number fits if it has been used already raise a used error
		defer singleProviderSession.lock.Unlock() // in case of an error we need to unlock the session as its currently locked.
		return nil, utils.LavaFormatWarning("Data Reliability Session was already used", DataReliabilitySessionAlreadyUsedError, nil)
	}

	return singleProviderSession, nil
}

func (psm *ProviderSessionManager) GetSession(address string, epoch uint64, sessionId uint64, relayNumber uint64) (*SingleProviderSession, error) {
	if !psm.IsValidEpoch(epoch) { // fast checking to see if epoch is even relevant
		utils.LavaFormatError("GetSession", InvalidEpochError, &map[string]string{"RequestedEpoch": strconv.FormatUint(epoch, 10), "blockedEpochHeight": strconv.FormatUint(psm.blockedEpochHeight, 10), "blockDistanceForEpochValidity": strconv.FormatUint(psm.blockDistanceForEpochValidity, 10)})
		return nil, InvalidEpochError
	}

	providerSessionWithConsumer, err := psm.IsActiveConsumer(epoch, address)
	if err != nil {
		return nil, err
	}

	return psm.getSingleSessionFromProviderSessionWithConsumer(providerSessionWithConsumer, sessionId, epoch, relayNumber)
}

func (psm *ProviderSessionManager) registerNewConsumer(consumerAddr string, epoch uint64, maxCuForConsumer uint64) (*ProviderSessionsWithConsumer, error) {
	psm.lock.Lock()
	defer psm.lock.Unlock()
	if !psm.IsValidEpoch(epoch) { // checking again because we are now locked and epoch cant change now.
		utils.LavaFormatError("getActiveConsumer", InvalidEpochError, &map[string]string{"RequestedEpoch": strconv.FormatUint(epoch, 10)})
		return nil, InvalidEpochError
	}

	mapOfProviderSessionsWithConsumer, foundEpochInMap := psm.sessionsWithAllConsumers[epoch]
	if !foundEpochInMap {
		mapOfProviderSessionsWithConsumer = make(map[string]*ProviderSessionsWithConsumer)
		psm.sessionsWithAllConsumers[epoch] = mapOfProviderSessionsWithConsumer
	}

	providerSessionWithConsumer, foundAddressInMap := mapOfProviderSessionsWithConsumer[consumerAddr]
	if !foundAddressInMap {
		providerSessionWithConsumer = NewProviderSessionsWithConsumer(consumerAddr, &ProviderSessionsEpochData{MaxComputeUnits: maxCuForConsumer}, notDataReliabilityPSWC)
		mapOfProviderSessionsWithConsumer[consumerAddr] = providerSessionWithConsumer
	}
	return providerSessionWithConsumer, nil
}

func (psm *ProviderSessionManager) RegisterProviderSessionWithConsumer(consumerAddress string, epoch uint64, sessionId uint64, relayNumber uint64, maxCuForConsumer uint64) (*SingleProviderSession, error) {
	providerSessionWithConsumer, err := psm.IsActiveConsumer(epoch, consumerAddress)
	if err != nil {
		if ConsumerNotRegisteredYet.Is(err) {
			providerSessionWithConsumer, err = psm.registerNewConsumer(consumerAddress, epoch, maxCuForConsumer)
			if err != nil {
				return nil, utils.LavaFormatError("RegisterProviderSessionWithConsumer Failed to registerNewSession", err, nil)
			}
		} else {
			return nil, utils.LavaFormatError("RegisterProviderSessionWithConsumer Failed", err, nil)
		}
		utils.LavaFormatDebug("provider registered consumer", &map[string]string{"consumer": consumerAddress, "epoch": strconv.FormatUint(epoch, 10)})
	}
	return psm.getSingleSessionFromProviderSessionWithConsumer(providerSessionWithConsumer, sessionId, epoch, relayNumber)
}

func (psm *ProviderSessionManager) getActiveConsumer(epoch uint64, address string) (providerSessionWithConsumer *ProviderSessionsWithConsumer, err error) {
	psm.lock.RLock()
	defer psm.lock.RUnlock()
	if !psm.IsValidEpoch(epoch) { // checking again because we are now locked and epoch cant change now.
		utils.LavaFormatError("getActiveConsumer", InvalidEpochError, &map[string]string{"RequestedEpoch": strconv.FormatUint(epoch, 10)})
		return nil, InvalidEpochError
	}
	if mapOfProviderSessionsWithConsumer, consumerFoundInEpoch := psm.sessionsWithAllConsumers[epoch]; consumerFoundInEpoch {
		if providerSessionWithConsumer, consumerAddressFound := mapOfProviderSessionsWithConsumer[address]; consumerAddressFound {
			if providerSessionWithConsumer.atomicReadConsumerBlocked() == blockListedConsumer { // we atomic read block listed so we dont need to lock the provider. (double lock is always a bad idea.)
				// consumer is blocked.
				utils.LavaFormatWarning("getActiveConsumer", ConsumerIsBlockListed, &map[string]string{"RequestedEpoch": strconv.FormatUint(epoch, 10), "ConsumerAddress": address})
				return nil, ConsumerIsBlockListed
			}
			return providerSessionWithConsumer, nil // no error
		}
	}
	return nil, ConsumerNotRegisteredYet
}

func (psm *ProviderSessionManager) getSessionFromAnActiveConsumer(providerSessionWithConsumer *ProviderSessionsWithConsumer, sessionId uint64, epoch uint64) (singleProviderSession *SingleProviderSession, err error) {
	session, err := providerSessionWithConsumer.GetExistingSession(sessionId)
	if err == nil {
		return session, nil
	} else if SessionDoesNotExist.Is(err) {
		// if we don't have a session we need to create a new one.
		return providerSessionWithConsumer.createNewSingleProviderSession(sessionId, epoch)
	} else {
		return nil, utils.LavaFormatError("could not get existing session", err, nil)
	}
}

func (psm *ProviderSessionManager) ReportConsumer() (address string, epoch uint64, err error) {
	return "", 0, nil // TBD
}

// OnSessionDone unlocks the session gracefully, this happens when session finished with an error
func (psm *ProviderSessionManager) OnSessionFailure(singleProviderSession *SingleProviderSession) (err error) {
	return singleProviderSession.onSessionFailure()
}

// OnSessionDone unlocks the session gracefully, this happens when session finished successfully
func (psm *ProviderSessionManager) OnSessionDone(singleProviderSession *SingleProviderSession) (err error) {
	return singleProviderSession.onSessionDone()
}

func (psm *ProviderSessionManager) RPCProviderEndpoint() *RPCProviderEndpoint {
	return psm.rpcProviderEndpoint
}

// on a new epoch we are cleaning stale provider data, also we are making sure consumers who are trying to use past data are not capable to
func (psm *ProviderSessionManager) UpdateEpoch(epoch uint64) {
	psm.lock.Lock()
	defer psm.lock.Unlock()
	if epoch > psm.blockDistanceForEpochValidity {
		psm.blockedEpochHeight = epoch - psm.blockDistanceForEpochValidity
	} else {
		psm.blockedEpochHeight = 0
	}
	psm.sessionsWithAllConsumers = filterOldEpochEntries(psm.blockedEpochHeight, psm.sessionsWithAllConsumers)
	psm.dataReliabilitySessionsWithAllConsumers = filterOldEpochEntries(psm.blockedEpochHeight, psm.dataReliabilitySessionsWithAllConsumers)
	// in the case of subscribe, we need to unsubscribe before deleting the key from storage.
	psm.subscriptionSessionsWithAllConsumers = psm.filterOldEpochEntriesSubscribe(psm.blockedEpochHeight, psm.subscriptionSessionsWithAllConsumers)
}

func (psm *ProviderSessionManager) filterOldEpochEntriesSubscribe(blockedEpochHeight uint64, allEpochsMap map[uint64]map[string]map[string]*RPCSubscription) map[uint64]map[string]map[string]*RPCSubscription {
	validEpochsMap := map[uint64]map[string]map[string]*RPCSubscription{}
	for epochStored, value := range allEpochsMap {
		if !IsEpochValidForUse(epochStored, blockedEpochHeight) {
			// epoch is not valid so we don't keep its key in the new map
			for _, consumers := range value { // unsubscribe
				for _, subscription := range consumers {
					if subscription.Sub == nil { // validate subscription not nil
						utils.LavaFormatError("filterOldEpochEntriesSubscribe Error", SubscriptionPointerIsNilError, &map[string]string{"subscripionId": subscription.Id})
					} else {
						subscription.Sub.Unsubscribe()
					}
				}
			}
			continue
		}
		// if epochStored is ok, copy the value into the new map
		validEpochsMap[epochStored] = value
	}
	return validEpochsMap
}

func filterOldEpochEntries[T any](blockedEpochHeight uint64, allEpochsMap map[uint64]T) (validEpochsMap map[uint64]T) {
	// In order to avoid running over the map twice, (1. mark 2. delete.) better technique is to copy and filter
	// which has better O(n) vs O(2n)
	validEpochsMap = map[uint64]T{}
	for epochStored, value := range allEpochsMap {
		if !IsEpochValidForUse(epochStored, blockedEpochHeight) {
			// epoch is not valid so we don't keep its key in the new map
			continue
		}
		// if epochStored is ok, copy the value into the new map
		validEpochsMap[epochStored] = value
	}
	return
}

func (psm *ProviderSessionManager) ProcessUnsubscribe(apiName string, subscriptionID string, consumerAddress string, epoch uint64) error {
	psm.lock.Lock()
	defer psm.lock.Unlock()
	mapOfConsumers, foundMapOfConsumers := psm.subscriptionSessionsWithAllConsumers[epoch]
	if !foundMapOfConsumers {
		return utils.LavaFormatError("Couldn't find epoch in psm.subscriptionSessionsWithAllConsumers", nil, &map[string]string{"epoch": strconv.FormatUint(epoch, 10), "address": consumerAddress})
	}
	mapOfSubscriptionId, foundMapOfSubscriptionId := mapOfConsumers[consumerAddress]
	if !foundMapOfSubscriptionId {
		return utils.LavaFormatError("Couldn't find consumer address in psm.subscriptionSessionsWithAllConsumers", nil, &map[string]string{"epoch": strconv.FormatUint(epoch, 10), "address": consumerAddress})
	}

	var err error
	if apiName == TendermintUnsubscribeAll {
		// unsubscribe all subscriptions
		for _, v := range mapOfSubscriptionId {
			if v.Sub == nil {
				err = utils.LavaFormatError("ProcessUnsubscribe TendermintUnsubscribeAll mapOfSubscriptionId Error", SubscriptionPointerIsNilError, &map[string]string{"subscripionId": subscriptionID})
			} else {
				v.Sub.Unsubscribe()
			}
		}
		psm.subscriptionSessionsWithAllConsumers[epoch][consumerAddress] = make(map[string]*RPCSubscription) // delete the entire map.
		return err
	}

	subscription, foundSubscription := mapOfSubscriptionId[subscriptionID]
	if !foundSubscription {
		return utils.LavaFormatError("Couldn't find subscription Id in psm.subscriptionSessionsWithAllConsumers", nil, &map[string]string{"epoch": strconv.FormatUint(epoch, 10), "address": consumerAddress, "subscriptionId": subscriptionID})
	}

	if subscription.Sub == nil {
		err = utils.LavaFormatError("ProcessUnsubscribe Error", SubscriptionPointerIsNilError, &map[string]string{"subscripionId": subscriptionID})
	} else {
		subscription.Sub.Unsubscribe()
	}
	delete(psm.subscriptionSessionsWithAllConsumers[epoch][consumerAddress], subscriptionID) // delete subscription after finished with it
	return err
}

func (psm *ProviderSessionManager) addSubscriptionToStorage(subscription *RPCSubscription, consumerAddress string, epoch uint64) error {
	psm.lock.Lock()
	defer psm.lock.Unlock()
	// we already validated the epoch is valid in the GetSession no need to verify again.
	_, foundEpoch := psm.subscriptionSessionsWithAllConsumers[epoch]
	if !foundEpoch {
		// this is the first time we subscribe in this epoch
		psm.subscriptionSessionsWithAllConsumers[epoch] = make(map[string]map[string]*RPCSubscription)
	}

	_, foundSubscriptions := psm.subscriptionSessionsWithAllConsumers[epoch][consumerAddress]
	if !foundSubscriptions {
		// this is the first subscription added in this epoch. we need to create the map
		psm.subscriptionSessionsWithAllConsumers[epoch][consumerAddress] = make(map[string]*RPCSubscription)
	}

	_, foundSubscription := psm.subscriptionSessionsWithAllConsumers[epoch][consumerAddress][subscription.Id]
	if !foundSubscription {
		// we shouldnt find a subscription already in the storage.
		psm.subscriptionSessionsWithAllConsumers[epoch][consumerAddress][subscription.Id] = subscription
		return nil // successfully added subscription to storage
	}

	// if we get here we found a subscription already in the storage and we need to return an error as we can't add two subscriptions with the same id
	return utils.LavaFormatError("addSubscription", SubscriptionAlreadyExistsError, &map[string]string{"SubscriptionId": subscription.Id, "epoch": strconv.FormatUint(epoch, 10), "address": consumerAddress})
}

func (psm *ProviderSessionManager) ReleaseSessionAndCreateSubscription(session *SingleProviderSession, subscription *RPCSubscription, consumerAddress string, epoch uint64) error {
	err := psm.OnSessionDone(session)
	if err != nil {
		return utils.LavaFormatError("Failed ReleaseSessionAndCreateSubscription", err, nil)
	}
	return psm.addSubscriptionToStorage(subscription, consumerAddress, epoch)
}

// try to disconnect the subscription incase we got an error.
// if fails to find assumes it was unsubscribed normally
func (psm *ProviderSessionManager) SubscriptionEnded(consumerAddress string, epoch uint64, subscriptionID string) {
	psm.lock.Lock()
	defer psm.lock.Unlock()
	mapOfConsumers, foundMapOfConsumers := psm.subscriptionSessionsWithAllConsumers[epoch]
	if !foundMapOfConsumers {
		return
	}
	mapOfSubscriptionId, foundMapOfSubscriptionId := mapOfConsumers[consumerAddress]
	if !foundMapOfSubscriptionId {
		return
	}

	subscription, foundSubscription := mapOfSubscriptionId[subscriptionID]
	if !foundSubscription {
		return
	}

	if subscription.Sub == nil { // validate subscription not nil
		utils.LavaFormatError("SubscriptionEnded Error", SubscriptionPointerIsNilError, &map[string]string{"subscripionId": subscription.Id})
	} else {
		subscription.Sub.Unsubscribe()
	}
	delete(psm.subscriptionSessionsWithAllConsumers[epoch][consumerAddress], subscriptionID) // delete subscription after finished with it
}

// Called when the reward server has information on a higher cu proof and usage and this providerSessionsManager needs to sync up on it
func (psm *ProviderSessionManager) UpdateSessionCU(consumerAddress string, epoch uint64, sessionID uint64, newCU uint64) error {
	// load the session and update the CU inside
	psm.lock.Lock()
	defer psm.lock.Unlock()
	if !psm.IsValidEpoch(epoch) { // checking again because we are now locked and epoch cant change now.
		return utils.LavaFormatError("UpdateSessionCU", InvalidEpochError, &map[string]string{"RequestedEpoch": strconv.FormatUint(epoch, 10)})
	}

	providerSessionsWithConsumerMap, ok := psm.sessionsWithAllConsumers[epoch]
	if !ok {
		return utils.LavaFormatError("UpdateSessionCU Failed", EpochIsNotRegisteredError, &map[string]string{"epoch": strconv.FormatUint(epoch, 10)})
	}
	providerSessionWithConsumer, foundConsumer := providerSessionsWithConsumerMap[consumerAddress]
	if !foundConsumer {
		return utils.LavaFormatError("UpdateSessionCU Failed", ConsumerIsNotRegisteredError, &map[string]string{"epoch": strconv.FormatUint(epoch, 10), "consumer": consumerAddress})
	}

	usedCu := providerSessionWithConsumer.atomicReadUsedComputeUnits() // check used cu now
	if usedCu < newCU {
		// if newCU proof is higher than current state, update.
		providerSessionWithConsumer.atomicWriteUsedComputeUnits(newCU)
	}
	return nil
}

// Returning a new provider session manager
func NewProviderSessionManager(rpcProviderEndpoint *RPCProviderEndpoint, numberOfBlocksKeptInMemory uint64) *ProviderSessionManager {
	return &ProviderSessionManager{
		rpcProviderEndpoint:                     rpcProviderEndpoint,
		blockDistanceForEpochValidity:           numberOfBlocksKeptInMemory,
		sessionsWithAllConsumers:                map[uint64]map[string]*ProviderSessionsWithConsumer{},
		dataReliabilitySessionsWithAllConsumers: map[uint64]map[string]*ProviderSessionsWithConsumer{},
		subscriptionSessionsWithAllConsumers:    map[uint64]map[string]map[string]*RPCSubscription{},
	}
}

func IsEpochValidForUse(targetEpoch uint64, blockedEpochHeight uint64) bool {
	return targetEpoch > blockedEpochHeight
}
