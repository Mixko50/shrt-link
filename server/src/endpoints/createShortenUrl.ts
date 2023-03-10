import { Elysia } from "elysia";
import { Shrt } from "../models/shrtTable";
import { ShrtRepository } from "../repository/shrt";
import { UrlRequest } from "../types/request";

import { BasedResponseElysia } from "../types/response";
import { generateSlug } from "../utils/generateSlug";

export const createShortenUrlController = (app: Elysia) => app
    .post('/create', async ({ body: { full_url, slug }, set }) => {
        const shrtRepo = new ShrtRepository();

        if (slug) {
            // Check duplicate link
            const check = await shrtRepo.getUrlBySlug(slug)

            if (check) {
                set.status = 400
                return {
                    success: false,
                    data: {
                        error_message: "Duplicate slug"
                    }
                }
            }

            // Create short link
            const create = await Shrt.create({ long_url: full_url, slug: slug });

            return {
                success: true,
                data: {
                    long_url: create.long_url,
                    slug: create.slug
                }
            }
        } else {
            const generatedSlug: string = generateSlug(6)
            const create = await Shrt.create({ long_url: full_url, slug: generatedSlug });

            return {
                success: true,
                data: {
                    long_url: create.long_url,
                    slug: create.slug
                }
            }
        }
    }, {
        schema: {
            body: UrlRequest,
            response: BasedResponseElysia
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
