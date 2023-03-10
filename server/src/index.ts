import { Elysia } from 'elysia'
import { errorHandling } from './errors/errorHandling'
import { connectDatabase } from './repository/connect'
import { route } from './routes/apiRoute'

// Check env
if (!process.env.DATABASE_URL) throw Error("Database_url not found")
if (!process.env.AUTO_MIGRATE) throw Error("AUTO_MIGRATE config not found")

// Initialize database
await connectDatabase()

// Elysia
const app = new Elysia()
    .use(errorHandling)
    .use(route)
    .listen(process.env.PORT ?? 8080)

console.log(`🦊 Elysia is running at http://localhost:${app.server?.port}`)

