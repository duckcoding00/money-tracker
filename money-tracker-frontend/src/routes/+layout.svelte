<script lang="ts">
	import '../app.css';
	import { onMount, type Snippet } from 'svelte';
	import type { LayoutData } from './$types';
	import AOS from 'aos';
	import 'aos/dist/aos.css';

	onMount(() => {
		AOS.init({
			duration: 800,
			once: true,
			offset: 120,
			easing: 'ease-out-cubic'
		});
	});

	let { data, children }: { data: LayoutData; children: Snippet } = $props();
</script>

{@render children()}
