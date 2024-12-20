<script lang="ts">
	import * as Alert from '$lib/components/ui/alert';
	import { Button } from '$lib/components/ui/button';
	import * as Card from '$lib/components/ui/card';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import * as Tooltip from '$lib/components/ui/tooltip/index';
	import { Info, Loader2 } from 'lucide-svelte';
	import type { Snippet } from 'svelte';
	import FormField from './FormField.svelte';

	type TunnelFormProps = {
		name: string;
		username: string;
		host: string;
		target: string;
		srcPort: number;
		destPort: number;
		handleSubmit: (e: SubmitEvent) => void;
	};

	let {
		name = $bindable(),
		username = $bindable(),
		host = $bindable(),
		target = $bindable(),
		srcPort = $bindable(),
		destPort = $bindable(),
		handleSubmit
	}: TunnelFormProps = $props();

	let errors: Record<string, string> = $state({});
	let isSubmitting = $state(false);

	function validateForm() {
		const newErrors: Record<string, string> = {};

		if (!name?.trim()) {
			newErrors.name = 'Name is required';
		}

		if (!username?.trim()) {
			newErrors.username = 'Username is required';
		}

		if (!host?.trim()) {
			newErrors.host = 'Host is required';
		}

		if (!target?.trim()) {
			newErrors.target = 'Target is required';
		}

		if (!srcPort) {
			newErrors.srcPort = 'Source port is required';
		} else if (srcPort < 1 || srcPort > 65535) {
			newErrors.srcPort = 'Port must be between 1 and 65535';
		}

		if (!destPort) {
			newErrors.destPort = 'Destination port is required';
		} else if (destPort < 1 || destPort > 65535) {
			newErrors.destPort = 'Port must be between 1 and 65535';
		}

		errors = newErrors;
		return Object.keys(newErrors).length === 0;
	}

	async function onSubmit(e: SubmitEvent) {
		e.preventDefault();
		if (!validateForm()) return;

		isSubmitting = true;
		try {
			await handleSubmit(e);
		} catch (error) {
			errors.submit = error instanceof Error ? error.message : 'An unknown error occurred';
		} finally {
			isSubmitting = false;
		}
	}
</script>

{#snippet formField(children: Snippet, label: string, tooltip: string)}
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
{/snippet}

<Card.Root class="mx-auto w-full max-w-lg">
	<Card.Header>
		<Card.Title>Configure SSH Tunnel</Card.Title>
	</Card.Header>
	<Card.Content>
		<form onsubmit={onSubmit} class="space-y-6">
			<div class="grid gap-6">
				<FormField label="Name" tooltip="A descriptive name for this tunnel" {errors}>
					<Input
						type="text"
						id="name"
						bind:value={name}
						class={errors.name ? 'border-red-500' : ''}
					/>
				</FormField>
				<div class="grid grid-cols-2 gap-4">
					<FormField label="Username" tooltip="Your SSH username" {errors}>
						<Input
							type="text"
							id="username"
							bind:value={username}
							class={errors.username ? 'border-red-500' : ''}
						/>
					</FormField>

					<FormField label="Host" tooltip="The host to connect to" {errors}>
						<Input
							type="text"
							id="host"
							bind:value={host}
							class={errors.host ? 'border-red-500' : ''}
						/>
					</FormField>
				</div>

				<FormField label="Target" tooltip="The target to connect to" {errors}>
					<Input
						type="text"
						id="target"
						bind:value={target}
						class={errors.target ? 'border-red-500' : ''}
					/>
				</FormField>

				<div class="grid grid-cols-2 gap-4">
					<FormField label="Local Port" tooltip="The local port (1-65535)" {errors}>
						<Input
							type="number"
							id="srcPort"
							bind:value={srcPort}
							class={errors.srcPort ? 'border-red-500' : ''}
						/>
					</FormField>

					<FormField label="Remote Port" tooltip="The remote port (1-65535)" {errors}>
						<Input
							type="number"
							id="destPort"
							bind:value={destPort}
							class={errors.destPort ? 'border-red-500' : ''}
						/>
					</FormField>
				</div>
			</div>

			{#if errors.submit}
				<Alert.Root variant="destructive">
					<Alert.Description>{errors.submit}</Alert.Description>
				</Alert.Root>
			{/if}

			<Button type="submit" class="w-full" disabled={isSubmitting}>
				{#if isSubmitting}
					<Loader2 class="mr-2 h-4 w-4 animate-spin" />
					Configuring...
				{:else}
					Configure Tunnel
				{/if}
			</Button>
		</form>
	</Card.Content>
</Card.Root>
