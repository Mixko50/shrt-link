import Elysia from "elysia";

export const errorHandling = (app: Elysia) => app
    .onError(({ code, error, set }) => {
        if (code === 'NOT_FOUND') {
            set.status = 404
            return {
                success: false,
                data: {
                    error_message: error.message,
                }
            }
        }
    })