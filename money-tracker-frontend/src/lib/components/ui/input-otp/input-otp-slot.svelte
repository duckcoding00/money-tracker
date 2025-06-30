<script lang="ts">
	import { PinInput as InputOTPPrimitive } from 'bits-ui';
	import { cn } from '$lib/utils.js';

	let {
		ref = $bindable(null),
		cell,
		class: className,
		...restProps
	}: InputOTPPrimitive.CellProps = $props();
</script>

<InputOTPPrimitive.Cell
	{cell}
	bind:ref
	data-slot="input-otp-slot"
	class={cn(
		'relative flex size-12 items-center justify-center border-4 border-black bg-white text-lg font-black tracking-wider text-black uppercase transition-all outline-none first:border-l last:border-r hover:shadow-[4px_4px_0px_0px_#666666]',
		cell.isActive && 'z-10 border-black bg-gray-100 shadow-[4px_4px_0px_0px_#000000]',
		className
	)}
	{...restProps}
>
	{cell.char}
	{#if cell.hasFakeCaret}
		<div class="pointer-events-none absolute inset-0 flex items-center justify-center">
			<div class="animate-caret-blink h-5 w-1 bg-black duration-1000"></div>
		</div>
	{/if}
</InputOTPPrimitive.Cell>
