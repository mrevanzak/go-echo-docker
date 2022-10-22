### Middleware
Middleware adalah proses perantara dari request ke response. Middleware yang saya pelajari adalah middleware untuk logging, middleware untuk authentication, dan middleware untuk authorization. Selain yang saya sebutkan itu terdapat pula syntax middleware lainnya seperti Echo.Pre(), Echo.Use(), dll.

#### Logging
Logging merupakan laporan dari suatu proses yang terjadi. Logging dapat dilakukan dengan menggunakan middleware. Pada contoh ini saya menggunakan middleware untuk logging dengan menggunakan middleware dari bawaan Echo yaitu Echo.LoggerWithConfig() dan menyertakan custom log.

### Authentication
Authentication adalah proses untuk mengidentifikasi user yang sedang login. Pada contoh ini saya menggunakan middleware untuk authentication dengan menggunakan middleware dari JWT.