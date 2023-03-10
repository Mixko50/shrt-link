export class BadRequestError extends Error {
    constructor(message?: string) {
        super(message || 'Bad request')
        Object.setPrototypeOf(this, BadRequestError.prototype)
        this.name = "Bad request"
    }
}