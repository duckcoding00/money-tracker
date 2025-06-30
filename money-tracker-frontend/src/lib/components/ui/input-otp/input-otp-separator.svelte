<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';
	import type { WithElementRef } from '$lib/utils.js';
	import DotIcon from '@lucide/svelte/icons/dot';

	let {
		ref = $bindable(null),
		children,
		...restProps
	}: WithElementRef<HTMLAttributes<HTMLDivElement>> = $props();
</script>

<div
	bind:this={ref}
	data-slot="input-otp-separator"
	role="separator"
	class="mx-2 text-2xl font-black text-black"
	{...restProps}
>
	{#if children}
		{@render children?.()}
	{:else}
		<DotIcon class="h-6 w-6 fill-black stroke-black" />
	{/if}
</div>
