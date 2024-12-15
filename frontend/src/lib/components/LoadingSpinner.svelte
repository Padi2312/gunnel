<script lang="ts">
	type LoadingSpinnerProps = {
		size?: 'xs' | 'sm' | 'md' | 'lg' | 'xl';
	};
	let { size = 'md' }: LoadingSpinnerProps = $props();

	const sizes = {
		xs: 'h-4 w-4',
		sm: 'h-8 w-8',
		md: 'h-16 w-16',
		lg: 'h-24 w-24',
		xl: 'h-32 w-32'
	};
</script>

<div class="flex items-center justify-center">
	<div
		class={`animate-spin ${sizes[size]} border-border rounded-full border-4 border-solid border-t-transparent`}
	></div>
</div>
