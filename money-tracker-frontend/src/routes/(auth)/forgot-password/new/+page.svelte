<script lang="ts">
	import { enhance } from '$app/forms';
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import Button from '$lib/components/ui/button/button.svelte';
	import Input from '$lib/components/ui/input/input.svelte';
	import Label from '$lib/components/ui/label/label.svelte';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
	let username = $derived(page.url.searchParams.get('username'));
	let token = $derived(page.url.searchParams.get('token'));

	let successAlert = $state('');
	let errorAlert = $state<string[]>([]);
	let loading = $state(false);
	let password = $state('');
	let confirm_password = $state('');

	function passwordEnhance() {
		loading = true;
		return async ({ result }: any) => {
			try {
				if (result.type === 'redirect') {
					successAlert = 'Your new password is set! Please log in and enjoy!';
					setTimeout(() => {
						goto(result.location, {
							invalidateAll: true
						});
					}, 2000);
				} else if (result.type === 'failure') {
					successAlert = '';
					let { details, message } = result.data;

					const errors: string[] = [];

					if (details) {
						if (typeof details === 'object' && details !== null) {
							Object.keys(details).forEach((field) => {
								if (details[field]) {
									errors.push(details[field]);
								}
							});
						} else if (typeof details === 'string') {
							errors.push(details);
						}
					} else if (message) {
						errors.push(message);
					}

					errorAlert = errors;
				}
			} catch (error) {
				console.error('error registerEnhance', error);
				errorAlert = ['An unexpected error occurred. Please try again.'];
				successAlert = '';
			} finally {
				loading = false;
			}
		};
	}
</script>

<div class="flex min-h-screen items-center justify-center bg-gray-50 px-4 py-8">
	<div class="w-full max-w-3xl">
		<div
			class={`grid gap-4 transition-all duration-300 ${successAlert || (errorAlert && errorAlert.length > 0) ? 'lg:grid-cols-[1fr_300px]' : 'grid-cols-1'}`}
		>
			<!-- Main Card -->
			<div
				class={`flex transition-all duration-300 ${successAlert || (errorAlert && errorAlert.length > 0) ? 'lg:justify-end' : 'lg:justify-center'}`}
			>
				<div class="w-full max-w-md">
					<div
						class="rotate-1 transform border-8 border-black bg-white p-8 shadow-[20px_20px_0px_0px_#666666]"
					>
						<div class="mb-8 text-center">
							<div
								class="mb-4 inline-block rotate-2 transform border-4 border-black bg-black px-4 py-2 shadow-[6px_6px_0px_0px_#666666]"
							>
								<h1 class="text-2xl font-black tracking-wider text-white uppercase">
									🔑 RESET PASSWORD
								</h1>
							</div>
							<p class="text-lg font-black tracking-wider text-black uppercase">
								Set Your New Password
							</p>
							<p class="mt-2 text-sm font-black tracking-wide text-gray-700 uppercase">
								{username ? `For: ${username}` : 'Please check your email link'}
							</p>
						</div>

						<form action="?/update" method="post" class="space-y-6" use:enhance={passwordEnhance}>
							<input type="hidden" name="username" value={username} />
							<input type="hidden" name="token" value={token} />

							<div class="space-y-3">
								<Label
									for="password"
									class="text-lg font-black tracking-wider text-black uppercase"
								>
									New Password
								</Label>
								<Input
									id="password"
									name="password"
									type="password"
									placeholder="ENTER NEW PASSWORD"
									class="h-12 rotate-1 transform border-4 border-black bg-white font-bold tracking-wide text-black uppercase shadow-[4px_4px_0px_0px_#666666] transition-all duration-200 placeholder:text-gray-500 focus:rotate-0 focus:border-gray-800 focus:ring-0"
									bind:value={password}
									required
									autocomplete="new-password"
								/>
								<p
									class="-rotate-1 transform text-xs font-bold tracking-wide text-gray-600 uppercase"
								>
									KEEP IT SECRET!
								</p>
							</div>

							<div class="space-y-3">
								<Label
									for="confirm_password"
									class="text-lg font-black tracking-wider text-black uppercase"
								>
									Confirm Password
								</Label>
								<Input
									id="confirm_password"
									name="confirm_password"
									type="password"
									placeholder="CONFIRM NEW PASSWORD"
									class="h-12 -rotate-1 transform border-4 border-black bg-white font-bold tracking-wide text-black uppercase shadow-[4px_4px_0px_0px_#666666] transition-all duration-200 placeholder:text-gray-500 focus:rotate-0 focus:border-gray-800 focus:ring-0"
									bind:value={confirm_password}
									required
									autocomplete="new-password"
								/>
								<p
									class="rotate-1 transform text-xs font-bold tracking-wide text-gray-600 uppercase"
								>
									MAKE SURE IT MATCHES!
								</p>
							</div>

							<div class="flex flex-col gap-4 pt-4">
								<Button
									class="h-14 -rotate-1 transform border-6 border-black bg-black text-lg font-black tracking-widest text-white uppercase shadow-[8px_8px_0px_0px_#666666] transition-all duration-200 hover:translate-x-2 hover:translate-y-2 hover:rotate-0 hover:bg-gray-800 hover:shadow-[4px_4px_0px_0px_#333333]"
									type="submit"
									disabled={loading}
								>
									{#if loading}
										⏳ RESETTING...
									{:else}
										🚀 RESET PASSWORD
									{/if}
								</Button>

								<Button
									class="h-12 rotate-1 transform border-4 border-black bg-white font-black tracking-wide text-black uppercase shadow-[6px_6px_0px_0px_#666666] transition-all duration-200 hover:translate-x-1 hover:translate-y-1 hover:rotate-0 hover:bg-gray-100 hover:shadow-[3px_3px_0px_0px_#333333]"
									variant="outline"
									type="button"
									onclick={() => {
										password = '';
										confirm_password = '';
									}}
								>
									🔄 CLEAR FORM
								</Button>
							</div>
						</form>
					</div>
				</div>
			</div>

			<!-- Alert Column -->
			{#if successAlert || (errorAlert && errorAlert.length > 0)}
				<div class="flex flex-col justify-center">
					{#if successAlert}
						{@render successMessage()}
					{:else if errorAlert && errorAlert.length > 0}
						{@render errorMessage()}
					{/if}
				</div>
			{/if}
		</div>
	</div>
</div>

{#snippet successMessage()}
	<div
		class="animate-in slide-in-from-right-5 max-w-xs rotate-1 transform border-4 border-black bg-gray-800 p-4 shadow-[8px_8px_0px_0px_#666666] duration-300"
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
					<p class="text-sm font-bold tracking-wide text-gray-300 uppercase">
						{successAlert}
					</p>
				</div>
			</div>
		</div>
	</div>
{/snippet}

{#snippet errorMessage()}
	{#if errorAlert && errorAlert.length > 0}
		<div
			class="animate-in slide-in-from-right-5 max-w-xs -rotate-1 transform border-4 border-black bg-red-600 p-4 shadow-[8px_8px_0px_0px_#666666] duration-300"
		>
			<div class="flex items-start">
				<div class="flex-shrink-0">
					<div class="flex h-6 w-6 items-center justify-center border-2 border-black bg-white">
						<span class="text-xs font-black text-red-600">!</span>
					</div>
				</div>
				<div class="ml-4">
					<h3 class="text-lg font-black tracking-wide text-white uppercase">ERROR!</h3>
					<div class="mt-2">
						{#each errorAlert as error}
							<p class="text-sm font-bold tracking-wide text-white uppercase">{error}</p>
						{/each}
					</div>
				</div>
			</div>
		</div>
	{/if}
{/snippet}
