import Elysia from "elysia";
import { createShortenUrlController } from "../endpoints/createShortenUrl";
import { redirect } from "../endpoints/redirect";
import { retrieveOriginalUrl } from "../endpoints/retrieveOriginal";

export const route = (app: Elysia) => app
    .group("/api", app => app
        .use(createShortenUrlController)
        .use(retrieveOriginalUrl)
    )
    .use(redirect)