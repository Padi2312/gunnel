<script lang="ts">
	import { onMount } from 'svelte';
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import { GetConfig, EditTunnel } from '../../../lib/wailsjs/go/internal/Store';
	import type { internal } from '../../../lib/wailsjs/go/models';
	import { goto } from '$app/navigation';

	let tunnel: internal.Tunnel = $state({
		id: '',
		name: '',
		username: '',
		host: '',
		target: '',
		srcPort: undefined as unknown as number,
		destPort: undefined as unknown as number
	});

	onMount(async () => {
		const config = await GetConfig();
		const tunnelId = $page.params.id;
		const existingTunnel = config.tunnels.find((t) => t.id === tunnelId);
		if (existingTunnel) {
			tunnel = { ...existingTunnel };
		}
	});

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
			await EditTunnel(tunnel);
			goto('/');
		} catch (error) {
			console.error(error);
			alert('Failed to edit tunnel');
		}
	};
</script>

<div class="container mx-auto">
	<h1 class="mb-4 text-xl">Edit Tunnel</h1>
	<form onsubmit={handleSubmit} class="space-y-4">
		<div class="space-y-2">
			<Label for="name">Name</Label>
			<Input
				type="text"
				id="name"
				bind:value={tunnel.name}
				placeholder="Enter tunnel name"
				required
			/>
		</div>

		<div class="space-y-2">
			<Label for="username">Username</Label>
			<Input
				type="text"
				id="username"
				bind:value={tunnel.username}
				placeholder="Enter username"
				required
			/>
		</div>

		<div class="space-y-2">
			<Label for="host">Host</Label>
			<Input type="text" id="host" bind:value={tunnel.host} placeholder="Enter host" required />
		</div>

		<div class="space-y-2">
			<Label for="target">Target</Label>
			<Input
				type="text"
				id="target"
				bind:value={tunnel.target}
				placeholder="Enter target"
				required
			/>
		</div>

		<div class="space-y-2">
			<Label for="sourcePort">Local Port</Label>
			<Input
				type="number"
				id="sourcePort"
				bind:value={tunnel.srcPort}
				placeholder="Enter source port"
				required
			/>
		</div>

		<div class="space-y-2">
			<Label for="destinationPort">Remote Port</Label>
			<Input
				type="number"
				id="destinationPort"
				bind:value={tunnel.destPort}
				placeholder="Enter destination port"
				required
			/>
		</div>

		<Button type="submit" class="w-full">Submit</Button>
	</form>
</div>
