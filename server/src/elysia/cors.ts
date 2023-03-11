import cors from "@elysiajs/cors";
import Elysia from "elysia";

export const corsConfig = (app: Elysia) => app
    .use(cors({
        origin: true,
        methods: ['GET', 'POST']
    }))