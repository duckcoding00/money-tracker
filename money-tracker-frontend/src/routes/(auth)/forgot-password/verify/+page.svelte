<script lang="ts">
	import { enhance } from '$app/forms';
	import { goto } from '$app/navigation';
	import { page } from '$app/state';
	import Button from '$lib/components/ui/button/button.svelte';
	import InputOtpGroup from '$lib/components/ui/input-otp/input-otp-group.svelte';
	import InputOtpSlot from '$lib/components/ui/input-otp/input-otp-slot.svelte';
	import InputOtp from '$lib/components/ui/input-otp/input-otp.svelte';
	import Input from '$lib/components/ui/input/input.svelte';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();

	let username = $derived(page.url.searchParams.get('username'));
	let message = $state('');
	let error = $state('');
	let loading = $state(false);
	let otp = $state('');

	function verifyEnhance() {
		loading = true;
		return async ({ result }: any) => {
			console.log(result);
			try {
				if (result.type === 'redirect') {
					message = 'Verification complete! Let’s set your new password';
					setTimeout(() => {
						goto(result.location);
					}, 1000);
				} else if (result.type === 'failure') {
					let { message } = result.data;
					error = message;
				}
			} catch (error) {
				console.error('error verifyEnhance', error);
			} finally {
				loading = false;
			}
		};
	}
</script>

<div class="flex min-h-screen items-center justify-center bg-gray-50 px-4 py-8">
	<div class="relative w-full max-w-4xl">
		<div
			class={`grid gap-4 transition-all duration-300 ${message || error ? 'grid-cols-1 lg:grid-cols-[1fr_280px]' : 'grid-cols-1'}`}
		>
			<!-- Main card -->
			<div
				class={`flex justify-center transition-all duration-300 ${message || error ? 'lg:justify-end' : 'lg:justify-center'}`}
			>
				<div class="w-full max-w-md">
					<div
						class="rotate-1 transform border-8 border-black bg-white p-8 shadow-[20px_20px_0px_0px_#666666]"
					>
						<div class="mb-8 text-center">
							<div
								class="mb-4 inline-block rotate-2 transform border-4 border-black bg-black px-4 py-2 shadow-[6px_6px_0px_0px_#666666]"
							>
								<h1 class="text-2xl font-black tracking-wider text-white uppercase">🔐 RESET</h1>
							</div>
							<p class="text-lg font-black tracking-wider text-black uppercase">
								ONE STEP MORE, JUST
							</p>
							<p class="mt-2 text-sm font-black tracking-wide text-gray-700 uppercase">
								📧 ENTER YOUR OTP BELOW
							</p>
							<p class="mt-2 text-sm font-black tracking-wide text-gray-700 uppercase">
								Check in your email, or spam email
							</p>
						</div>

						<form action="?/verify" method="post" class="space-y-8" use:enhance={verifyEnhance}>
							<div class="text-center">
								<InputOtp maxlength={6} class="justify-center" bind:value={otp} name="otp">
									{#snippet children({ cells })}
										<InputOtpGroup>
											{#each cells as cell (cell)}
												<InputOtpSlot {cell} />
											{/each}
										</InputOtpGroup>
									{/snippet}
								</InputOtp>

								<input name="username" type="hidden" bind:value={username} />
							</div>

							<div class="space-y-4">
								<Button
									class="h-12 w-full -rotate-1 transform border-6 border-black bg-black text-lg font-black tracking-widest text-white uppercase shadow-[10px_10px_0px_0px_#666666] transition-all duration-200 hover:translate-x-2 hover:translate-y-2 hover:rotate-0 hover:bg-gray-800 hover:shadow-[5px_5px_0px_0px_#333333]"
									type="submit"
									disabled={loading}
								>
									{#if loading}
										⏳ VERIFYING...
									{:else}
										VERIFY
									{/if}
								</Button>

								<div class="flex gap-3">
									<Button
										class="h-10 flex-1 rotate-1 transform border-4 border-black bg-white text-sm font-black tracking-wide text-black uppercase shadow-[6px_6px_0px_0px_#666666] transition-all duration-200 hover:translate-x-1 hover:translate-y-1 hover:rotate-0 hover:bg-gray-100 hover:shadow-[3px_3px_0px_0px_#333333]"
										variant="outline"
									>
										📤 RESEND
									</Button>

									<Button
										class="h-10 flex-1 -rotate-1 transform border-4 border-black bg-white text-sm font-black tracking-wide text-black uppercase shadow-[6px_6px_0px_0px_#666666] transition-all duration-200 hover:translate-x-1 hover:translate-y-1 hover:rotate-0 hover:bg-gray-100 hover:shadow-[3px_3px_0px_0px_#333333]"
										variant="outline"
										onclick={() => {
											otp = '';
										}}
									>
										🔄 RESET
									</Button>
								</div>
							</div>
						</form>
					</div>
				</div>
			</div>

			<!-- Alert column -->
			<div class="flex flex-col justify-center">
				{#if message}
					{@render successMessage()}
				{:else if error}
					{@render errorMessage()}
				{/if}
			</div>
		</div>
	</div>
</div>

{#snippet successMessage()}
	<div
		class="animate-in slide-in-from-right-5 rotate-1 transform border-4 border-black bg-gray-800 p-4 shadow-[8px_8px_0px_0px_#666666] duration-300"
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
					<p class="text-sm font-bold tracking-wide text-gray-300 uppercase">{message}</p>
				</div>
			</div>
		</div>
	</div>
{/snippet}
{#snippet errorMessage()}
	<div
		class="animate-in slide-in-from-right-5 -rotate-1 transform border-4 border-black bg-red-600 p-4 shadow-[8px_8px_0px_0px_#666666] duration-300"
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
					<p class="text-sm font-bold tracking-wide text-white uppercase">{error}</p>
				</div>
			</div>
		</div>
	</div>
{/snippet}
