// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	compatibilityDate: '2025-07-15',
	devtools: { enabled: true },

	modules: [ '@pinia/nuxt' ],

	runtimeConfig:
	{
		public: { api: process.env.NUXT_PUBLIC_API_HOST }
	},

	ssr: false,

	css: ['@/assets/styles/index.scss'],

	vite:
	{
		css:
		{
			preprocessorOptions:
			{
				scss:
				{
					silenceDeprecations: ['import', 'global-builtin', 'legacy-js-api'],
					additionalData: `@use "@/assets/styles/base/_variables.scss" as *;  @use "@/assets/styles/base/_mixins.scss" as *; @use "@/assets/styles/base/_normalize.scss";`,
				},
			},
		},
	},
})
