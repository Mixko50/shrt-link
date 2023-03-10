const s = "abcdefghijklmnopqrstuvwxyz0123456789";
export const generateSlug = (len: number): string => {
    return Array(len + 1).join().split('').map(() => { return s.charAt(Math.floor(Math.random() * s.length)); }).join('');
}