import cors from "@elysiajs/cors";
import Elysia from "elysia";

export const corsConfig = (app: Elysia) => app
    .use(cors({
        origin: /m.mixkomii.com$/,
        methods: ['GET', 'POST'],
        allowedHeaders: ['Content-Type', 'application/json'],
        exposedHeaders: ['Content-Type', 'application/json']
    }))