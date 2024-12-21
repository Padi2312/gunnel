<script lang="ts">
	import { goto } from '$app/navigation';
	import * as Popover from '$lib/components/ui/popover';
	import { internal } from '$lib/wailsjs/go/models';
	import EllipsisVertical from 'lucide-svelte/icons/ellipsis-vertical';
	import Pencil from 'lucide-svelte/icons/pencil';
	import Trash from 'lucide-svelte/icons/trash';

	type TunnelMenu = {
		tunnel: internal.Tunnel;
		onDelete: (tunnel: internal.Tunnel) => void;
	};
	let { tunnel, onDelete }: TunnelMenu = $props();
</script>

<Popover.Root>
	<Popover.Trigger class="outline-none ring-0">
		<EllipsisVertical />
	</Popover.Trigger>
	<Popover.Content class="!p-2" side="left" align="start">
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
