<script lang="ts">
	import TunnelMenu from '$lib/components/gunnel/TunnelMenu.svelte';
	import LoadingSpinner from '$lib/components/LoadingSpinner.svelte';
	import { Button } from '$lib/components/ui/button';
	import { Card } from '$lib/components/ui/card';
	import type { internal } from '$lib/wailsjs/go/models';

	type TunnelCardProps = {
		tunnel: internal.Tunnel;
		isActive: boolean;
		loading: boolean;
		onConnect: (tunnel: internal.Tunnel) => void;
		onDisconnect: (tunnel: internal.Tunnel) => void;
		onDelete: (tunnel: internal.Tunnel) => void;
	};
	let { tunnel, isActive, loading, onConnect, onDisconnect, onDelete }: TunnelCardProps = $props();
</script>

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
		{#if isActive}
			<Button variant="destructive" onclick={() => onDisconnect(tunnel)}>Disconnect</Button>
		{:else}
			<Button onclick={() => onConnect(tunnel)} disabled={loading}>
				{#if loading}
					<LoadingSpinner size="xs" />
					Connecting...
				{:else}
					Connect
				{/if}
			</Button>
		{/if}
		<TunnelMenu {tunnel} {onDelete} />
	</div>
</Card>
