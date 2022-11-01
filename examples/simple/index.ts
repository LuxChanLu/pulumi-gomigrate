import * as gomigrate from "@pulumi/gomigrate";

const page = new gomigrate.StaticPage("page", {
    indexContent: "<html><body><p>Hello world!</p></body></html>",
});

export const bucket = page.bucket;
export const url = page.websiteUrl;
