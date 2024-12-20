<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { Card } from '$lib/components/ui/card';
	import { onMount } from 'svelte';
	import { Connect, Disconnect, GetActiveTunnels } from '../lib/wailsjs/go/internal/Handler';
	import { GetConfig, RemoveTunnel } from '../lib/wailsjs/go/internal/Store';
	import type { internal } from '../lib/wailsjs/go/models';
	import LoadingSpinner from '../lib/components/LoadingSpinner.svelte';
	import { goto } from '$app/navigation';
	import * as Popover from '$lib/components/ui/popover/index';
	import EllipsisVertical from 'lucide-svelte/icons/ellipsis-vertical';
	import Trash from 'lucide-svelte/icons/trash';
	import Pencil from 'lucide-svelte/icons/pencil';

	let config: internal.Config = $state({
		sshPrivateKeyPath: '',
		username: '',
		tunnels: []
	} as unknown as internal.Config);

	let activeTunnelIds: string[] = $state([]);
	let loadingTunnels: string[] = $state([]);

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
</script>

<!-- <div class="container mx-auto p-4"> -->
<div class="container mx-auto h-full space-y-2">
	{#each getTunnels() as tunnel (tunnel.id)}
		<Card class="flex items-center justify-between p-4">
			<div class="flex items-center gap-4">
				<div class="flex-1">
					<h3 class="text-lg font-semibold">{tunnel.name}</h3>
					<div class="text-muted-foreground mt-2 flex flex-col space-y-1 text-xs">
						<div class="flex space-x-4">
							<div>
								<span class="font-medium">Host:</span>
								{tunnel.host}
							</div>
							<div>
								<span class="font-medium">Target:</span>
								{tunnel.target}
							</div>
						</div>
						<div class="flex space-x-4">
							<div>
								<span class="font-medium">Local Port:</span>
								{tunnel.srcPort}
							</div>
							<div>
								<span class="font-medium">Remote Port:</span>
								{tunnel.destPort}
							</div>
						</div>
					</div>
				</div>
			</div>
			<div class="flex items-center gap-4">
				{#if isTunnelActive(tunnel)}
					<Button variant="destructive" onclick={() => onDisconnect(tunnel)}>Disconnect</Button>
				{:else}
					<Button onclick={() => onConnect(tunnel)} disabled={loadingTunnels.includes(tunnel.id)}>
						{#if loadingTunnels.includes(tunnel.id)}
							<LoadingSpinner size="xs" />
							Connecting...
						{:else}
							Connect
						{/if}
					</Button>
				{/if}

				<Popover.Root>
					<Popover.Trigger class="outline-none ring-0">
						<EllipsisVertical />
					</Popover.Trigger>
					<Popover.Content class="!p-2">
						<div class="flex flex-col space-y-2 p-0">
							<button
								class="hover:bg-primary-foreground flex items-center space-x-2 rounded p-2 text-left text-sm outline-none ring-0"
								onclick={() => goto(`/edit?id=${tunnel.id}`)}
							>
								<Pencil class="h-3 w-3" />
								<span>Edit</span>
							</button>
							<button
								class="hover:bg-primary-foreground flex items-center space-x-2 rounded p-2 text-left text-sm text-red-500 outline-none ring-0"
								onclick={() => onDelete(tunnel)}
							>
								<Trash class="h-3 w-3" />
								<span>Delete</span>
							</button>
						</div>
					</Popover.Content>
				</Popover.Root>
			</div>
		</Card>
	{:else}
		<div class="flex flex-col items-center justify-center h-full w-full space-y-6 py-12">
			<div class="flex flex-col items-center space-y-2">
				<h2 class="text-2xl font-semibold">No Tunnels Found</h2>
				<p class="text-muted-foreground text-center">Create your first SSH tunnel to get started</p>
			</div>
			<Button class="px-6" onclick={() => goto('/add')}>
				<span class="mr-2">+</span>
				Add New Tunnel
			</Button>
		</div>
	{/each}
</div>
<!-- </div> -->
