import { t } from 'elysia'

export const UrlRequest = t.Object({
    full_url: t.String(t.RegEx(/((?:(?:http?|ftp)[s]*:\/\/)?[a-z0-9-%\/\&=?\.]+\.[a-z]{2,4}\/?([^\s<>\#%"\,\{\}\\|\\\^\[\]`]+)?)/gi)), // eslint-disable-line
    slug: t.Optional(t.String())
})

export const RetrieveUrlRequest = t.Object({
    shrt_url: t.String(t.RegEx(/^(https|http):\/\/(m.mixkomii.com|localhost)/))
})