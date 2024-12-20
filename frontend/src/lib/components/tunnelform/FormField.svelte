<script lang="ts">
	import type { Snippet } from 'svelte';
	import * as Tooltip from '$lib/components/ui/tooltip';
	import { Label } from '$lib/components/ui/label';
	import Info from 'lucide-svelte/icons/info';

	type FormFieldProps = {
		children: Snippet;
		label: string;
		tooltip: string;
		errors: Record<string, string>;
	};
	let { label, tooltip, errors, children }: FormFieldProps = $props();
</script>

<div class="space-y-2">
	<Label for={label} class="flex items-center space-x-2">
		<span>{label}</span>
		<Tooltip.Provider>
			<Tooltip.Root>
				<Tooltip.Trigger>
					<Info class="h-4 w-4 text-gray-500" />
				</Tooltip.Trigger>
				<Tooltip.Content>
					<p class="text-sm">{tooltip}</p>
				</Tooltip.Content>
			</Tooltip.Root>
		</Tooltip.Provider>
	</Label>
	{@render children()}
	{#if errors.name}<span class="text-xs text-red-500">{errors.name}</span>{/if}
</div>
