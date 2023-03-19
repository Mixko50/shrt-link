import { Elysia } from 'elysia'
import { corsConfig } from './elysia/cors'
import { errorHandling } from './elysia/errorHandling'
import { connectDatabase } from './repository/connect'
import { route } from './routes/apiRoute'

// Check env
if (!process.env.DATABASE_URL) throw Error("Database_url not found")
if (!process.env.AUTO_MIGRATE) throw Error("AUTO_MIGRATE config not found")

// Initialize database
await connectDatabase()

// Elysia
const app = new Elysia()
    .use(corsConfig)
    .use(errorHandling)
    .use(route)
    .listen(process.env.PORT ?? 8080)

console.log(`💧 Shrt server is running at http://localhost:${app.server?.port}`)

