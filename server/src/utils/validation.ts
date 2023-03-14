export const isAlphanumeric = (text: string): boolean => {
    const regex = new RegExp(/^[a-z0-9]+$/i)
    return regex.test(text)
}