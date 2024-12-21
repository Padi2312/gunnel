<script lang="ts">
	import { Card, CardContent, CardHeader, CardTitle } from '$lib/components/ui/card';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import {
		GetConfig,
		SetUsername
	} from '$lib/wailsjs/go/internal/Store';
	import type { internal } from '$lib/wailsjs/go/models';
	import { onMount } from 'svelte';

	let config: internal.Config = {
		sshPrivateKeyPath: '',
		username: '',
		tunnels: []
	} as unknown as internal.Config;

	onMount(async () => {
		const loadedConfig = await GetConfig();
		config.sshPrivateKeyPath = loadedConfig.sshPrivateKeyPath;
		config.username = loadedConfig.username;
	});

	const onChangeUsername = (e: Event) => {
		config.username = (e.target as HTMLInputElement).value;
		SetUsername(config.username);
	};
</script>

<div class="container mx-auto">
	<Card class="w-full">
		<CardHeader>
			<CardTitle>Settings</CardTitle>
		</CardHeader>
		<CardContent>
			<form class="space-y-4">
				<div class="space-y-2">
					<Label for="username">Username</Label>
					<Input
						type="text"
						id="username"
						bind:value={config.username}
						placeholder="Enter your username"
						required
						onchange={onChangeUsername}
					/>
					<span class="text-muted-foreground text-xs">
						Your SSH username for remote connections
					</span>
				</div>
			</form>
		</CardContent>
	</Card>
</div>
