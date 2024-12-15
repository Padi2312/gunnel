<script lang="ts">
	import { goto } from '$app/navigation';
	import * as Sidebar from '$lib/components/ui/sidebar/index';
	import Home from 'lucide-svelte/icons/home';
	import Plus from 'lucide-svelte/icons/plus';
	import Cog from 'lucide-svelte/icons/cog';
</script>

<Sidebar.Root>
	<Sidebar.Header>
		<h1 class="text-sm font-bold">Gunnel</h1>
	</Sidebar.Header>
	<Sidebar.Content>
		<Sidebar.Group />
		<Sidebar.MenuButton onclick={() => goto('/')}>
			<Home />
			<span> Home </span>
		</Sidebar.MenuButton>
		<Sidebar.MenuButton onclick={() => goto('/add')}>
			<Plus />
			<span> Add </span>
		</Sidebar.MenuButton>
		<Sidebar.Group />
	</Sidebar.Content>
	<Sidebar.Footer>
		<Sidebar.MenuButton onclick={() => goto('/settings')}>
			<Cog />
			<span> Settings </span>
		</Sidebar.MenuButton>
	</Sidebar.Footer>
</Sidebar.Root>
