// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./migration";
export * from "./migrations";
export * from "./provider";

// Import resources to register:
import { Migration } from "./migration";
import { Migrations } from "./migrations";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "gomigrate:index:Migration":
                return new Migration(name, <any>undefined, { urn })
            case "gomigrate:index:Migrations":
                return new Migrations(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("gomigrate", "index", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("gomigrate", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:gomigrate") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
