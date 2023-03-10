import { t } from "elysia"

export const ErrorResponseElysia = t.Object({
    error_message: t.String(),
    detail: t.Optional(t.String())
})

export const ShrtResponseElysia = t.Object({
    long_url: t.Optional(t.String()),
    slug: t.Optional(t.String()),
})

export const BasedResponseElysia = t.Object({
    success: t.Boolean(),
    data: ShrtResponseElysia
})

export const ShrtResponse = t.Object({
    success: t.Boolean(),
    data: t.Object({
        long_url: t.Optional(t.String()),
        slug: t.Optional(t.String()),
        error_message: t.Optional(t.String()),
        detail: t.Optional(t.String())
    })
})