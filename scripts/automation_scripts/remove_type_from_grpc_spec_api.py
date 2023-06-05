import json

input_files = [
    "/home/user/go/src/lava/cookbook/specs/spec_add_cosmoshub.json",
    "/home/user/go/src/lava/cookbook/specs/spec_add_juno.json",
    "/home/user/go/src/lava/cookbook/specs/spec_add_lava.json",
    "/home/user/go/src/lava/cookbook/specs/spec_add_osmosis.json",
]

for f in input_files:
    with open(f,"r") as fr:
        data = json.load(fr)
    for spec in data['proposal']['specs']:
        for api in spec["apis"]:
            for apiInterface in api["api_interfaces"]:
                if apiInterface["interface"] == "grpc":
                    apiInterface["type"] = ""

    with open(f,"w") as fr:
        json.dump(data,fr,indent=4)