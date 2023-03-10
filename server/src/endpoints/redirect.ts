import Elysia, { t } from "elysia";
import { ShrtRepository } from "../repository/shrt";

export const redirect = (app: Elysia) => app
    .get('/:id', async ({ params: { id }, set }) => {
        const shrtRepo = new ShrtRepository();

        // Check duplicate link
        const check = await shrtRepo.getUrlBySlug(id)

        if (check) {
            // Redirect
            set.redirect = check.long_url

            // Update visit
            await shrtRepo.updateVisit(id)

        } else {
            set.redirect = process.env.BASE_URL
        }
    }, {
        schema: {
            params: t.Object({
                id: t.String({
                    minLength: 1
                })
            })
        }
    }).onError(({ set }) => {
        set.redirect = process.env.BASE_URL
    })