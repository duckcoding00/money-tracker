<script lang="ts">
	import Navbar from '$lib/components/custom/Navbar.svelte';

	let { children } = $props();
</script>

<Navbar />
{@render children()}
