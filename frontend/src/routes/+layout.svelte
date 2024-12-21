<script lang="ts">
	import * as Sidebar from '$lib/components/ui/sidebar/index.js';
	import AppSidebar from '$lib/components/app-sidebar.svelte';
	import '../app.css';

	let { children } = $props();
</script>

<Sidebar.Provider>
	<AppSidebar />
	<main class="h-screen w-screen">
		<div class="absolute">
			<Sidebar.Trigger />
		</div>
		<div class="h-full w-full p-2">
			{@render children?.()}
		</div>
	</main>
</Sidebar.Provider>
