export default defineNuxtConfig({
    runtimeConfig: {
        isServer: true,
        public: {
            siteUrl: process.env.VITE_SITE_URL,
            baseUrl: process.env.VITE_API_DOMAIN,
        }
    },
    css: [
        "reset-css/reset.css",
        'animate.css/animate.min.css'
    ]
})
