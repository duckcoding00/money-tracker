<script lang="ts">
	import { enhance } from '$app/forms';
	import type { SubmitFunction } from '@sveltejs/kit';
	import Button from '../ui/button/button.svelte';
	import { goto } from '$app/navigation';

	type props = {
		action?: string;
		enhance?: SubmitFunction;
		isActive?: boolean;
	};

	let { action, enhance: logoutEnhance, isActive = false }: props = $props();
	function logoutHandle() {}
</script>

<div
	class="flex min-h-screen w-55 flex-col justify-between gap-4 border-r-8 border-black bg-gray-200 px-4 py-6"
>
	<div class="space-y-8">
		<div
			class="rotate-1 transform border-4 border-black bg-white p-4 shadow-[6px_6px_0px_0px_#666666]"
		>
			<p class="text-sm font-black tracking-wider text-black uppercase">MAIN MENU</p>
			<div class="mt-4 space-y-3">
				<button
					class="w-full -rotate-1 transform border-4 border-black bg-gray-800 px-4 py-2 text-left shadow-[4px_4px_0px_0px_#333333] transition-all duration-200 hover:translate-x-1 hover:translate-y-1 hover:shadow-[2px_2px_0px_0px_#111111]"
					onclick={() => {
						goto('/dashboard');
					}}
				>
					<span class="text-xs font-black tracking-wide text-white uppercase">📊 DASHBOARD</span>
				</button>
				<button
					class="w-full rotate-1 transform border-4 border-black px-4 py-2 text-left shadow-[4px_4px_0px_0px_#666666] transition-all duration-200 hover:translate-x-1 hover:translate-y-1 hover:shadow-[2px_2px_0px_0px_#333333] {!isActive ? 'bg-gray-400 cursor-not-allowed opacity-50' : 'bg-white hover:bg-gray-100'}"
					onclick={() => {
						if (isActive) {
							goto('/activity');
						}
					}}
					disabled={!isActive}
				>
					<span class="text-xs font-black tracking-wide {!isActive ? 'text-gray-600' : 'text-black'} uppercase">⚡ ACTIVITY</span>
				</button>
				<button
					class="w-full -rotate-1 transform border-4 border-black px-4 py-2 text-left shadow-[4px_4px_0px_0px_#333333] transition-all duration-200 hover:translate-x-1 hover:translate-y-1 hover:shadow-[2px_2px_0px_0px_#111111] {!isActive ? 'bg-gray-500 cursor-not-allowed opacity-50' : 'bg-gray-600'}"
					onclick={() => {
						if (isActive) {
							goto('/profile');
						}
					}}
					disabled={!isActive}
				>
					<span class="text-xs font-black tracking-wide text-white uppercase">👤 PROFILE</span>
				</button>
			</div>
		</div>
	</div>

	<div class="mb-16">
		<div
			class="-rotate-1 transform border-4 border-black bg-black p-4 shadow-[6px_6px_0px_0px_#666666]"
		>
			<p class="text-sm font-black tracking-wider text-white uppercase">SECOND MENU</p>
			<div class="mt-4 space-y-3">
				<button
					class="w-full rotate-1 transform border-4 border-white px-4 py-2 text-left shadow-[4px_4px_0px_0px_#333333] transition-all duration-200 hover:translate-x-1 hover:translate-y-1 hover:shadow-[2px_2px_0px_0px_#666666] {!isActive ? 'bg-gray-400 cursor-not-allowed opacity-50' : 'bg-white hover:bg-gray-100'}"
					onclick={() => {
						if (isActive) {
							goto('/settings');
						}
					}}
					disabled={!isActive}
				>
					<span class="text-xs font-black tracking-wide {!isActive ? 'text-gray-600' : 'text-black'} uppercase">⚙️ SETTINGS</span>
				</button>

				<form {action} method="post" use:enhance={logoutEnhance}>
					<button
						type="submit"
						class="w-full -rotate-1 transform border-4 border-white bg-gray-800 px-4 py-2 text-left shadow-[4px_4px_0px_0px_#333333] transition-all duration-200 hover:translate-x-1 hover:translate-y-1 hover:bg-gray-900 hover:shadow-[2px_2px_0px_0px_#111111]"
						onclick={() => {
							console.log('logout clicked');
						}}
					>
						<span class="text-xs font-black tracking-wide text-white uppercase">🚪 LOGOUT</span>
					</button>
				</form>
			</div>
		</div>
	</div>
</div>
