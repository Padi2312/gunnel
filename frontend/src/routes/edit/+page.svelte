<script lang="ts">
	import { goto } from '$app/navigation';
	import TunnelForm from '$lib/components/gunnel/tunnelform/TunnelForm.svelte';
	import { UpdateTunnel } from '$lib/wailsjs/go/internal/Store';
	import { internal } from '$lib/wailsjs/go/models';
	import type { PageData } from './$types';

	type PageProps = {
		data: PageData;
	};

	let { data }: PageProps = $props();

	let tunnel: internal.Tunnel = $state(data.tunnel);

	const handleSubmit = async () => {
		await UpdateTunnel(tunnel);
		await goto('/');
	};
</script>

<div class="container mx-auto">
	<h1 class="mb-4 text-xl">Edit Tunnel</h1>
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
