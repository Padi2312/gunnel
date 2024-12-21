<script lang="ts">
	import { goto } from '$app/navigation';
	import TunnelCard from '$lib/components/gunnel/TunnelCard.svelte';
	import { Button } from '$lib/components/ui/button';
	import { Connect, Disconnect, GetActiveTunnels } from '$lib/wailsjs/go/internal/Handler';
	import { GetConfig, RemoveTunnel } from '$lib/wailsjs/go/internal/Store';
	import type { internal } from '$lib/wailsjs/go/models';
	import { LayoutGrid, List } from 'lucide-svelte';
	import { onMount } from 'svelte';

	let config: internal.Config = $state({
		sshPrivateKeyPath: '',
		username: '',
		tunnels: []
	} as unknown as internal.Config);

	let activeTunnelIds: string[] = $state([]);
	let loadingTunnels: string[] = $state([]);
	let isGrouped = $state(false);

	onMount(() => {
		init();
	});

	const init = async () => {
		config = await GetConfig();
		const activeTunnelMap = await GetActiveTunnels();
		activeTunnelIds = activeTunnelMap;
	};

	const getTunnels = () => {
		// List connected tunnels first and then disconnected tunnels
		const activeTunnels = getActiveTunnels();
		const disconnectedTunnels = config.tunnels.filter((t) => !activeTunnelIds.includes(t.id));
		return [...activeTunnels, ...disconnectedTunnels];
	};

	const onConnect = async (tunnel: internal.Tunnel) => {
		if (!tunnel) {
			return;
		}
		loadingTunnels.push(tunnel.id);
		await Connect(tunnel, config.sshPrivateKeyPath);
		const activeTunnelMap = await GetActiveTunnels();
		activeTunnelIds = activeTunnelMap;
		loadingTunnels = loadingTunnels.filter((id) => id !== tunnel.id);
	};

	const onDisconnect = async (tunnel: internal.Tunnel) => {
		if (!tunnel) {
			return;
		}
		await Disconnect(tunnel);
		const activeTunnelMap = await GetActiveTunnels();
		activeTunnelIds = activeTunnelMap;
	};

	const onDelete = async (tunnel: internal.Tunnel) => {
		if (!tunnel) {
			return;
		}
		await RemoveTunnel(tunnel.id);
		init();
	};

	const isTunnelActive = (tunnel: internal.Tunnel) => {
		return activeTunnelIds.includes(tunnel.id);
	};

	const getActiveTunnels = (): internal.Tunnel[] => {
		let activeTunnels: internal.Tunnel[] = [];
		for (const tunnelId of activeTunnelIds) {
			const tunnel = config.tunnels.find((t) => t.id === tunnelId);
			if (tunnel) {
				activeTunnels.push(tunnel);
			}
		}
		return activeTunnels;
	};

	const groupTunnelsByHost = () => {
		const groupedTunnels: Record<string, internal.Tunnel[]> = {};
		for (const tunnel of config.tunnels) {
			if (!groupedTunnels[tunnel.host]) {
				groupedTunnels[tunnel.host] = [];
			}
			groupedTunnels[tunnel.host].push(tunnel);
		}
		return groupedTunnels;
	};
</script>

<div class="container mx-auto h-full space-y-2">
	<!-- Layout Toggle -->
	<div class="flex justify-end">
		<Button variant="ghost" size="sm" onclick={() => (isGrouped = !isGrouped)}>
			{#if isGrouped}
				<List class="mr-2 h-4 w-4" />
				List View
			{:else}
				<LayoutGrid class="mr-2 h-4 w-4" />
				Group by Server
			{/if}
		</Button>
	</div>

	{#if config.tunnels.length === 0}
		<div class="flex h-full w-full flex-col items-center justify-center space-y-6 py-12">
			<div class="flex flex-col items-center space-y-2">
				<h2 class="text-2xl font-semibold">No Tunnels Found</h2>
				<p class="text-muted-foreground text-center">Create your first SSH tunnel to get started</p>
			</div>
			<Button class="px-6" onclick={() => goto('/add')}>
				<span class="mr-2">+</span>
				Add New Tunnel
			</Button>
		</div>
	{/if}

	{#snippet iterateOverTunnels(tunnels: internal.Tunnel[])}
		{#each tunnels as tunnel (tunnel.id)}
			<TunnelCard
				{tunnel}
				isActive={isTunnelActive(tunnel)}
				loading={loadingTunnels.includes(tunnel.id)}
				{onConnect}
				{onDisconnect}
				{onDelete}
			/>
		{/each}
	{/snippet}

	{#if isGrouped}
		{#each Object.entries(groupTunnelsByHost()) as [host, tunnels] (host)}
			<div class="flex flex-col space-y-2 pb-4">
				<h2 class="text-lg font-semibold">{host}</h2>
				{@render iterateOverTunnels(tunnels)}
			</div>
		{/each}
	{:else}
		{@render iterateOverTunnels(config.tunnels)}
	{/if}
</div>
