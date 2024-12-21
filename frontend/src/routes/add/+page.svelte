<script lang="ts">
	import { onMount } from 'svelte';
	import { AddTunnel, GetConfig } from '$lib/wailsjs/go/internal/Store';
	import type { internal } from '$lib/wailsjs/go/models';
	import TunnelForm from '$lib/components/gunnel/tunnelform/TunnelForm.svelte';

	let tunnel: internal.Tunnel = $state({
		id: crypto.randomUUID(),
		name: '',
		username: '',
		host: '',
		target: '',
		srcPort: undefined as unknown as number,
		destPort: undefined as unknown as number
	});

	onMount(() => {
		trySetUsername();
	});

	const reset = () => {
		tunnel = {
			id: crypto.randomUUID(),
			name: '',
			username: '',
			host: '',
			target: '',
			srcPort: undefined as unknown as number,
			destPort: undefined as unknown as number
		};
	};

	const trySetUsername = async () => {
		const config = await GetConfig();
		tunnel.username = config.username;
	};

	const handleSubmit = async (e: SubmitEvent) => {
		e.preventDefault();
		if (
			!tunnel.name ||
			!tunnel.username ||
			!tunnel.host ||
			!tunnel.target ||
			!tunnel.srcPort ||
			!tunnel.destPort
		) {
			return;
		}
		try {
			await AddTunnel(tunnel);
			reset();
		} catch (error) {
			console.error(error);
			alert('Failed to add tunnel');
		}
	};
</script>

<div class="container mx-auto pb-4">
	<h1 class="mb-4 text-xl">Add New Tunnel</h1>
	<TunnelForm
		bind:name={tunnel.name}
		bind:username={tunnel.username}
		bind:host={tunnel.host}
		bind:target={tunnel.target}
		bind:srcPort={tunnel.srcPort}
		bind:destPort={tunnel.destPort}
		{handleSubmit}
	/>
</div>
