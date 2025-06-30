<script lang="ts">
	import type { SubmitFunction } from '@sveltejs/kit';
	import Button from '../ui/button/button.svelte';
	import Input from '../ui/input/input.svelte';
	import Label from '../ui/label/label.svelte';
	import { enhance } from '$app/forms';
	import { goto } from '$app/navigation';

	type props = {
		email?: string;
		password?: string;
		action?: string;
		loading?: boolean;
		successAlert?: string;
		errorAlert?: string[];
		enhance?: SubmitFunction;
		failedCredentials?: boolean;
	};

	let {
		email = $bindable(''),
		password = $bindable(''),
		action,
		loading = false,
		successAlert,
		errorAlert,
		enhance: loginEnhance,
		failedCredentials = false
	}: props = $props();

	function cancelForm() {
		email = '';
		password = '';
	}
</script>

<div class="rotate-1 transform">
	<form {action} method="post" class="space-y-6" use:enhance={loginEnhance}>
		{#if errorAlert}
			{@render errorMessage()}
		{/if}
		{#if successAlert}
			{@render successMessage()}
		{/if}

		<div class="space-y-3">
			<Label for="email" class="text-lg font-black tracking-wider text-black uppercase"
				>YOUR EMAIL</Label
			>
			<Input
				id="email"
				name="email"
				type="email"
				placeholder="ENTER YOUR EMAIL"
				class="h-12 -rotate-1 transform border-4 border-black bg-white font-bold tracking-wide text-black uppercase shadow-[4px_4px_0px_0px_#666666] transition-all duration-200 placeholder:text-gray-500 focus:rotate-0 focus:border-gray-800 focus:ring-0"
				bind:value={email}
			/>
			<p class="rotate-1 transform text-xs font-bold tracking-wide text-gray-600 uppercase">
				SECURE & ENCRYPTED!
			</p>
		</div>

		<div class="space-y-3">
			<Label for="password" class="text-lg font-black tracking-wider text-black uppercase"
				>YOUR PASSWORD</Label
			>
			<Input
				id="password"
				name="password"
				type="password"
				placeholder="ENTER YOUR PASSWORD"
				class="h-12 rotate-1 transform border-4 border-black bg-white font-bold tracking-wide text-black uppercase shadow-[4px_4px_0px_0px_#666666] transition-all duration-200 placeholder:text-gray-500 focus:rotate-0 focus:border-gray-800 focus:ring-0"
				bind:value={password}
			/>
			<p class="-rotate-1 transform text-xs font-bold tracking-wide text-gray-600 uppercase">
				KEEP IT SECRET!
			</p>
		</div>

		<div class="flex flex-col gap-4 pt-4">
			<Button
				class="h-14 -rotate-1 transform border-6 border-black bg-black text-lg font-black tracking-widest text-white uppercase shadow-[8px_8px_0px_0px_#666666] transition-all duration-200 hover:translate-x-2 hover:translate-y-2 hover:rotate-0 hover:bg-gray-800 hover:shadow-[4px_4px_0px_0px_#333333]"
				type="submit"
				disabled={loading}
			>
				{#if loading}
					⏳ LOGGING IN...
				{:else}
					🚀 LET'S DIVE IN!
				{/if}
			</Button>

			<Button
				class="h-12 rotate-1 transform border-4 border-black bg-white font-black tracking-wide text-black uppercase shadow-[6px_6px_0px_0px_#666666] transition-all duration-200 hover:translate-x-1 hover:translate-y-1 hover:rotate-0 hover:bg-gray-100 hover:shadow-[3px_3px_0px_0px_#333333]"
				variant="outline"
				onclick={() => cancelForm()}
			>
				🔄 START OVER
			</Button>
		</div>
	</form>
</div>

{#snippet errorMessage()}
	{#if errorAlert && errorAlert.length > 0}
		<div
			class="-rotate-1 transform border-4 border-black bg-white p-4 shadow-[8px_8px_0px_0px_#666666]"
		>
			<div class="flex items-start">
				<div class="flex-shrink-0">
					<div class="flex h-6 w-6 items-center justify-center border-2 border-black bg-black">
						<span class="text-xs font-black text-white">❌</span>
					</div>
				</div>
				<div class="ml-4">
					<h3 class="text-lg font-black tracking-wide text-black uppercase">FIX THESE ERRORS:</h3>
					<div class="mt-2">
						<ul class="space-y-2">
							{#each errorAlert as error}
								<li class="flex items-center gap-2">
									<div class="h-2 w-2 bg-black"></div>
									<span class="text-sm font-bold tracking-wide text-black uppercase">{error}</span>
								</li>
							{/each}
						</ul>
						{#if failedCredentials}
							<div class="mt-4">
								<Button
									class="h-10 -rotate-1 transform border-4 border-black bg-gray-800 px-4 py-2 font-black tracking-wide text-white uppercase shadow-[6px_6px_0px_0px_#666666] transition-all duration-200 hover:translate-x-1 hover:translate-y-1 hover:rotate-0 hover:bg-black hover:shadow-[3px_3px_0px_0px_#333333]"
									onclick={() => {
										goto('/forgot-password');
									}}
								>
									🔑 RESET PASSWORD
								</Button>
							</div>
						{/if}
					</div>
				</div>
			</div>
		</div>
	{/if}
{/snippet}

{#snippet successMessage()}
	{#if successAlert}
		<div
			class="rotate-1 transform border-4 border-black bg-gray-800 p-4 shadow-[8px_8px_0px_0px_#666666]"
		>
			<div class="flex items-start">
				<div class="flex-shrink-0">
					<div class="flex h-6 w-6 items-center justify-center border-2 border-white bg-white">
						<span class="text-xs font-black text-black">✓</span>
					</div>
				</div>
				<div class="ml-4">
					<h3 class="text-lg font-black tracking-wide text-white uppercase">SUCCESS!</h3>
					<div class="mt-2">
						<p class="text-sm font-bold tracking-wide text-gray-300 uppercase">{successAlert}</p>
					</div>
				</div>
			</div>
		</div>
	{/if}
{/snippet}
