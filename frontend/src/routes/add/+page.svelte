<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import * as Tooltip from '$lib/components/ui/tooltip/index';
	import Info from 'lucide-svelte/icons/info';
	import { onMount } from 'svelte';
	import { AddTunnel, GetConfig } from '../../lib/wailsjs/go/internal/Store';
	import type { internal } from '../../lib/wailsjs/go/models';

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

{#snippet labelTooltip(content: string)}
	<Tooltip.Provider>
		<Tooltip.Root>
			<Tooltip.Trigger>
				<Info class="h-4 w-4" />
			</Tooltip.Trigger>
			<Tooltip.Content>
				<p>{content}</p>
			</Tooltip.Content>
		</Tooltip.Root>
	</Tooltip.Provider>
{/snippet}

<div class="container mx-auto">
	<h1 class="mb-4 text-xl">Add New Tunnel</h1>
	<form onsubmit={handleSubmit} class="space-y-4">
		<div class="space-y-2">
			<Label for="name" class="flex items-center space-x-2">
				<span>Name</span>
				{@render labelTooltip('A descriptive name for this tunnel')}
			</Label>
			<Input
				type="text"
				id="name"
				bind:value={tunnel.name}
				placeholder="Enter tunnel name"
				required
			/>
		</div>

		<div class="space-y-2">
			<Label for="username" class="flex items-center space-x-2">
				<span>Username</span>
				{@render labelTooltip('Your SSH username for remote connections')}
			</Label>
			<Input
				type="text"
				id="username"
				bind:value={tunnel.username}
				placeholder="Enter username"
				required
			/>
		</div>

		<div class="space-y-2">
			<Label for="host" class="flex items-center space-x-2">
				<span>Host</span>
				{@render labelTooltip('The host to connect to')}
			</Label>
			<Input type="text" id="host" bind:value={tunnel.host} placeholder="Enter host" required />
		</div>

		<div class="space-y-2">
			<Label for="target" class="flex items-center space-x-2">
				<span>Target</span>
				{@render labelTooltip('The target to connect to')}
			</Label>
			<Input
				type="text"
				id="target"
				bind:value={tunnel.target}
				placeholder="Enter target"
				required
			/>
		</div>

		<div class="space-y-2">
			<Label for="sourcePort" class="flex items-center space-x-2">
				<span>Local Port</span>
				{@render labelTooltip('The local port to forward')}
			</Label>
			<Input
				type="number"
				id="sourcePort"
				bind:value={tunnel.srcPort}
				placeholder="Enter source port"
				required
			/>
		</div>

		<div class="space-y-2">
			<Label for="destinationPort" class="flex items-center space-x-2">
				<span>Remote Port</span>
				{@render labelTooltip('The remote port to forward')}
			</Label>
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
