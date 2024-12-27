INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `username`, `fullname`, `phone`, `password`, `role`, `active`)
VALUES (NULL, NULL, NULL, NULL, 'admin', 'Administrator', '089677271000', '$2a$14$3WPl7l1/ZnygAXV2BbJa1esKwJcGXTYEPh4KknYMt3.FVJxuF9cuy', 'admin', 1);
INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `username`, `fullname`, `phone`, `active`, `password`, `role`, `force_cpw`, `last_cpw`)
VALUES
	(2, NULL, '2024-07-22 05:36:47.679', NULL, 'customer', 'Customer', '08564763947', 1, '$2a$14$n1AOUkivJ3Ryi/fM8YKoOusHlq.4DZsX2F3cOYsvI.5Vq63AUCngG', 'customer', 0, '2024-07-12 06:56:37');

-- //pass: Merdeka1