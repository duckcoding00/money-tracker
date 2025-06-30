<script lang="ts">
	import { goto } from '$app/navigation';
	import Button from '$lib/components/ui/button/button.svelte';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
</script>

<section class="min-h-screen">
	<div
		class="mx-auto my-4 w-full max-w-7xl border-8 border-black bg-white shadow-[20px_20px_0px_0px_#000000]"
	>
		<div class="grid min-h-[80vh] grid-cols-12">
			<div class="col-span-8 border-r-8 border-black bg-gray-200">
				<div class="flex h-full flex-col justify-center p-12">
					<div class="mb-8">
						<h1
							class="-rotate-1 transform text-6xl leading-tight font-black tracking-wider text-black uppercase"
						>
							STOP WONDERING<br />
							WHERE YOUR<br />
							<span
								class="inline-block rotate-2 transform border-4 border-black bg-black px-4 py-2 text-white shadow-[8px_8px_0px_0px_#666666]"
							>
								MONEY GOES!
							</span>
						</h1>

						<div
							class="mt-8 rotate-1 transform border-4 border-black bg-white p-6 shadow-[10px_10px_0px_0px_#666666]"
						>
							<p class="text-xl font-bold tracking-wide text-black uppercase">
								TRACK EVERY PENNY → REACH EVERY GOAL<br />
								YOUR FINANCIAL FREEDOM STARTS HERE!
							</p>
						</div>
					</div>

					<div class="flex gap-6">
						<Button
							class="-rotate-1 transform border-6 border-black bg-black px-8 py-4 text-xl font-black tracking-wider text-white uppercase shadow-[12px_12px_0px_0px_#666666] transition-all duration-200 hover:translate-x-3 hover:translate-y-3 hover:rotate-0 hover:bg-gray-800 hover:shadow-[6px_6px_0px_0px_#333333]"
							size="lg"
							onclick={() => {
								goto('/register');
							}}
						>
							🚀 START HERE!
						</Button>

						<Button
							class="rotate-1 transform border-6 border-black bg-white px-6 py-4 text-lg font-black tracking-wider text-black uppercase shadow-[10px_10px_0px_0px_#666666] transition-all duration-200 hover:translate-x-2 hover:translate-y-2 hover:rotate-0 hover:bg-gray-100 hover:shadow-[5px_5px_0px_0px_#333333]"
							size="lg"
							onclick={() => {
								goto('/login');
							}}
						>
							💰 ALREADY JOIN?
						</Button>
					</div>
				</div>
			</div>

			<div class="col-span-4 flex items-center justify-center bg-gray-300 p-8">
				<div
					class="rotate-3 transform border-6 border-black bg-white p-8 shadow-[15px_15px_0px_0px_#666666]"
				>
					<h2 class="mb-4 -rotate-2 transform text-3xl font-black text-black uppercase">
						MONEY STATS
					</h2>
					<div class="space-y-4">
						<div
							class="border-4 border-black bg-gray-800 p-4 text-white shadow-[6px_6px_0px_0px_#333333]"
						>
							<span class="text-2xl font-black">$12,345</span>
							<p class="font-bold text-gray-300 uppercase">TOTAL SAVED</p>
						</div>
						<div class="border-4 border-black bg-white p-4 shadow-[6px_6px_0px_0px_#333333]">
							<span class="text-2xl font-black text-black">187</span>
							<p class="font-bold text-gray-600 uppercase">EXPENSES TRACKED</p>
						</div>
						<div
							class="border-4 border-black bg-gray-600 p-4 text-white shadow-[6px_6px_0px_0px_#333333]"
						>
							<span class="text-2xl font-black">95%</span>
							<p class="font-bold text-gray-300 uppercase">BUDGET ON TRACK</p>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</section>
