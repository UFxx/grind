// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	compatibilityDate: '2025-07-15',
	devtools: { enabled: true },

	modules:
	[
		'@pinia/nuxt',
		'@vueuse/nuxt',
		'nuxt-countdown',
		'@vueuse/motion/nuxt'
	],

	runtimeConfig:
	{
		public:
		{
			api     : process.env.NUXT_PUBLIC_API_HOST,
			botName : process.env.NUXT_PUBLIC_BOT_NAME
		},
	},

	ssr: false,

	app:
	{
		head:
		{
			script:
			[
				{ src: "https://telegram.org/js/telegram-web-app.js", defer: true },
			]
		}
	},

	css: ['@/assets/styles/index.scss'],

	vite:
	{
		server: { allowedHosts: true },
		css:
		{
			preprocessorOptions:
			{
				scss:
				{
					silenceDeprecations : ['import', 'global-builtin', 'legacy-js-api'],
					additionalData      : `
						@use "@/assets/styles/base/_transitions.scss" as *;
						@use "@/assets/styles/base/_variables.scss" as *;
						@use "@/assets/styles/base/_mixins.scss" as *;
						@use "@/assets/styles/base/_normalize.scss";
						@use "@/assets/styles/base/_fonts.scss";
					`,
				},
			},
		},
	},

	nitro:
	{
		devProxy:
		{
			'/api':
			{
				target       : process.env.NUXT_PUBLIC_API_HOST,
				changeOrigin : true
			}
		}
	},
})
