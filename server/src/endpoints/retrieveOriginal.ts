import Elysia from "elysia";
import { ShrtRepository } from "../repository/shrt";
import { RetrieveUrlRequest } from "../types/request";
import { BasedResponseElysia } from "../types/response";
import { isAlphanumeric } from "../utils/validation";

export const retrieveOriginalUrl = (app: Elysia) => app
    .post("/retrieve", async ({ body: { shrt_url }, set }) => {

        const slug: string = shrt_url.substring(shrt_url.lastIndexOf('/') + 1)

        if (!isAlphanumeric(slug)) {
            set.status = 400
            return {
                success: false,
                error: {
                    error_message: "Validation failed",
                    detail: "The slug is invalid",
                }
            }
        }

        const shrtRepo = new ShrtRepository();

        // Check duplicate link
        const check = await shrtRepo.getUrlBySlug(slug)

        if (check) {
            return {
                success: true,
                data: {
                    long_url: check.long_url,
                    slug: check.slug
                }
            }
        } else {
            set.status = 400
            return {
                success: false,
                error: {
                    error_message: "URL Not found",
                    detail: "The slug is not found",
                }
            }
        }

    }, {
        schema: {
            body: RetrieveUrlRequest,
            response: BasedResponseElysia
        }
    }).onError(({ code, set }) => {
        if (code == 'VALIDATION') {
            set.status = 400
            return {
                success: false,
                error: {
                    error_message: "Validation failed",
                    detail: "The url is invalid",
                }
            }
        }
    })
