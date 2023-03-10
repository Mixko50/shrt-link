import Elysia from "elysia";
import { ShrtRepository } from "../repository/shrt";
import { RetrieveUrlRequest } from "../types/request";
import { ShrtResponse } from "../types/response";

export const retrieveOriginalUrl = (app: Elysia) => app
    .post("/retrieve", async ({ body: { shrt_url }, set }) => {

        const slug: string = shrt_url.substring(shrt_url.lastIndexOf('/') + 1)

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
                success: true,
                data: {
                    error_message: "URL Not found"
                }
            }
        }

    }, {
        schema: {
            body: RetrieveUrlRequest,
            response: ShrtResponse
        }
    }).onError(({ code, error, set }) => {
        if (code == 'VALIDATION') {
            set.status = 400
            return {
                success: false,
                data: {
                    error_message: error.message,
                    detail: "The url is invalid"
                }
            }
        }
    })
