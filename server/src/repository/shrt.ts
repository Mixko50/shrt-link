import { Shrt } from "../models/shrtTable"

export class ShrtRepository {
    async getUrlBySlug(slug: string): Promise<Shrt | null> {
        const link: Shrt | null = await Shrt.findOne({
            where: {
                slug: slug
            }
        })
        return link
    }

    async updateVisit(slug: string) {
        const link = await Shrt.increment({
            visit: 1
        }, {
            where: {
                slug: slug
            }
        })
        return link
    }
}