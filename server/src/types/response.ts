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
    data: t.Optional(ShrtResponseElysia),
    error: t.Optional(ErrorResponseElysia)
})