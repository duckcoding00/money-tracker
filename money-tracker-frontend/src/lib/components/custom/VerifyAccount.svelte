<script lang="ts">
	import type { SubmitFunction } from '@sveltejs/kit';
	import Button from '../ui/button/button.svelte';
	import InputOtpGroup from '../ui/input-otp/input-otp-group.svelte';
	import InputOtpSlot from '../ui/input-otp/input-otp-slot.svelte';
	import InputOtp from '../ui/input-otp/input-otp.svelte';
	import { enhance } from '$app/forms';

	type props = {
		otp?: string;
		username?: string;
		loading?: boolean;
		action?: string;
		enhace?: SubmitFunction;
		successAlert?: string;
		errorAlert?: string[];
	};

	let {
		otp = $bindable(''),
		username = $bindable(''),
		action,
		loading = false,
		enhace: verifyEnhance,
		successAlert,
		errorAlert
	}: props = $props();
</script>

<div class="flex min-h-screen items-center justify-center bg-gray-50 px-4 py-8">
	<div class="grid w-full max-w-6xl grid-cols-1 gap-8 lg:grid-cols-3">
		<div class="hidden lg:block"></div>

		<div class="w-full">
			<div
				class="rotate-1 transform border-8 border-black bg-white p-8 shadow-[20px_20px_0px_0px_#666666]"
			>
				<div class="mb-8 text-center">
					<div
						class="mb-4 inline-block rotate-2 transform border-4 border-black bg-black px-4 py-2 shadow-[6px_6px_0px_0px_#666666]"
					>
						<h1 class="text-2xl font-black tracking-wider text-white uppercase">🔐 VERIFY</h1>
					</div>
					<p class="text-lg font-black tracking-wider text-black uppercase">YOUR ACCOUNT</p>
					<p class="mt-2 text-sm font-black tracking-wide text-gray-700 uppercase">
						📧 CHECK YOUR EMAIL FOR OTP
					</p>
				</div>

				<form {action} method="post" class="space-y-8" use:enhance={verifyEnhance}>
					<div class="text-center">
						<p class="mb-4 text-sm font-black tracking-wide text-black uppercase">
							ENTER 6-DIGIT CODE
						</p>
						<InputOtp maxlength={6} class="justify-center" bind:value={otp} name="otp">
							{#snippet children({ cells })}
								<InputOtpGroup>
									{#each cells as cell (cell)}
										<InputOtpSlot {cell} />
									{/each}
								</InputOtpGroup>
							{/snippet}
						</InputOtp>
						<input type="hidden" name="username" bind:value={username} />
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
								✅ VERIFY CODE
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

			<div class="mt-6 text-center">
				<div
					class="inline-block -rotate-1 transform border-4 border-black bg-gray-200 px-4 py-2 shadow-[4px_4px_0px_0px_#666666]"
				>
					<p class="text-xs font-black tracking-wide text-black uppercase">
						💡 DIDN'T GET CODE? CHECK SPAM FOLDER
					</p>
				</div>
			</div>
		</div>

		<div class="flex flex-col justify-center space-y-4">
			{#if errorAlert && errorAlert.length > 0}
				{@render errorMessage()}
			{/if}

			{#if successAlert}
				{@render successMessage()}
			{/if}
		</div>
	</div>
</div>

{#snippet errorMessage()}
	{#if errorAlert && errorAlert.length > 0}
		<div
			class="animate-in slide-in-from-right-5 -rotate-1 transform border-4 border-black bg-white p-4 shadow-[8px_8px_0px_0px_#666666] duration-300"
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
					</div>
				</div>
			</div>
		</div>
	{/if}
{/snippet}

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
					<p class="text-sm font-bold tracking-wide text-gray-300 uppercase">{successAlert}</p>
				</div>
			</div>
		</div>
	</div>
{/snippet}
