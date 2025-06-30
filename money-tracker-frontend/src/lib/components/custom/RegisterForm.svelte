<script lang="ts">
	import type { SubmitFunction } from '@sveltejs/kit';
	import Button from '../ui/button/button.svelte';
	import Input from '../ui/input/input.svelte';
	import Label from '../ui/label/label.svelte';
	import { enhance } from '$app/forms';

	type props = {
		email?: string;
		username?: string;
		password?: string;
		confirm_password?: string;
		action?: string;
		loading?: boolean;
		successAlert?: string;
		errorAlert?: string[];
		enhance?: SubmitFunction;
	};

	let {
		email = $bindable(''),
		username = $bindable(''),
		password = $bindable(''),
		confirm_password = $bindable(''),
		action,
		loading = false,
		successAlert,
		errorAlert,
		enhance: registerEnhance
	}: props = $props();

	function cancelForm() {
		email = '';
		username = '';
		password = '';
		confirm_password = '';
	}
</script>

<div class="transform -rotate-1">
	<form {action} method="post" class="space-y-6" use:enhance={registerEnhance}>
		{#if errorAlert}
			{@render errorMessage()}
		{/if}
		{#if successAlert}
			{@render successMessage()}
		{/if}
		
		<div class="space-y-3">
			<Label for="email" class="text-lg font-black text-black uppercase tracking-wider">YOUR EMAIL</Label>
			<Input
				id="email"
				name="email"
				type="email"
				placeholder="ENTER YOUR EMAIL"
				class="h-12 border-4 border-black bg-white font-bold text-black placeholder:text-gray-500 focus:border-gray-800 focus:ring-0 shadow-[4px_4px_0px_0px_#666666] uppercase tracking-wide transform rotate-1 focus:rotate-0 transition-all duration-200"
				bind:value={email}
			/>
			<p class="text-xs font-bold text-gray-600 uppercase tracking-wide transform -rotate-1">WE'LL SEND YOU AMAZING UPDATES!</p>
		</div>

		<div class="space-y-3">
			<Label for="username" class="text-lg font-black text-black uppercase tracking-wider">YOUR USERNAME</Label>
			<Input
				id="username"
				name="username"
				type="text"
				placeholder="ENTER YOUR USERNAME"
				class="h-12 border-4 border-black bg-white font-bold text-black placeholder:text-gray-500 focus:border-gray-800 focus:ring-0 shadow-[4px_4px_0px_0px_#666666] uppercase tracking-wide transform -rotate-1 focus:rotate-0 transition-all duration-200"
				bind:value={username}
			/>
			<p class="text-xs font-bold text-gray-600 uppercase tracking-wide transform rotate-1">MAKE IT UNIQUELY YOU!</p>
		</div>

		<div class="space-y-3">
			<Label for="password" class="text-lg font-black text-black uppercase tracking-wider">SECRET PASSWORD</Label>
			<Input
				id="password"
				name="password"
				type="password"
				placeholder="ENTER YOUR PASSWORD"
				class="h-12 border-4 border-black bg-white font-bold text-black placeholder:text-gray-500 focus:border-gray-800 focus:ring-0 shadow-[4px_4px_0px_0px_#666666] uppercase tracking-wide transform rotate-1 focus:rotate-0 transition-all duration-200"
				bind:value={password}
			/>
			<p class="text-xs font-bold text-gray-600 uppercase tracking-wide transform -rotate-1">MAKE IT STRONG LIKE A SUPERHERO!</p>
		</div>

		<div class="space-y-3">
			<Label for="confirm_password" class="text-lg font-black text-black uppercase tracking-wider">CONFIRM PASSWORD</Label>
			<Input
				id="confirm_password"
				name="confirm_password"
				type="password"
				placeholder="CONFIRM YOUR PASSWORD"
				class="h-12 border-4 border-black bg-white font-bold text-black placeholder:text-gray-500 focus:border-gray-800 focus:ring-0 shadow-[4px_4px_0px_0px_#666666] uppercase tracking-wide transform -rotate-1 focus:rotate-0 transition-all duration-200"
				bind:value={confirm_password}
			/>
			<p class="text-xs font-bold text-gray-600 uppercase tracking-wide transform rotate-1">DOUBLE-CHECK FOR EXTRA AWESOMENESS!</p>
		</div>

		<div class="flex flex-col gap-4 pt-4">
			<Button 
				class="h-14 border-6 border-black bg-black font-black text-white uppercase tracking-widest shadow-[8px_8px_0px_0px_#666666] hover:translate-x-2 hover:translate-y-2 hover:shadow-[4px_4px_0px_0px_#333333] hover:bg-gray-800 transition-all duration-200 transform rotate-1 hover:rotate-0 text-lg" 
				type="submit"
				disabled={loading}
			>
				{#if loading}
					⏳ JOINING...
				{:else}
					🚀 JOIN THE ADVENTURE!
				{/if}
			</Button>
			
			<Button
				class="h-12 border-4 border-black bg-white font-black text-black uppercase tracking-wide shadow-[6px_6px_0px_0px_#666666] hover:translate-x-1 hover:translate-y-1 hover:shadow-[3px_3px_0px_0px_#333333] hover:bg-gray-100 transition-all duration-200 transform -rotate-1 hover:rotate-0"
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
        <div class="border-4 border-black bg-white p-4 shadow-[8px_8px_0px_0px_#666666] transform rotate-1">
            <div class="flex items-start">
                <div class="flex-shrink-0">
                    <div class="h-6 w-6 border-2 border-black bg-black flex items-center justify-center">
                        <span class="text-white font-black text-xs">❌</span>
                    </div>
                </div>
                <div class="ml-4">
                    <h3 class="text-lg font-black text-black uppercase tracking-wide">FIX THESE ERRORS:</h3>
                    <div class="mt-2">
                        <ul class="space-y-2">
                            {#each errorAlert as error}
                                <li class="flex items-center gap-2">
                                    <div class="h-2 w-2 bg-black"></div>
                                    <span class="font-bold text-black uppercase text-sm tracking-wide">{error}</span>
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
    {#if successAlert}
        <div class="border-4 border-black bg-gray-800 p-4 shadow-[8px_8px_0px_0px_#666666] transform -rotate-1">
            <div class="flex items-start">
                <div class="flex-shrink-0">
                    <div class="h-6 w-6 border-2 border-white bg-white flex items-center justify-center">
                        <span class="text-black font-black text-xs">✓</span>
                    </div>
                </div>
                <div class="ml-4">
                    <h3 class="text-lg font-black text-white uppercase tracking-wide">SUCCESS!</h3>
                    <div class="mt-2">
                        <p class="font-bold text-gray-300 uppercase text-sm tracking-wide">{successAlert}</p>
                    </div>
                </div>
            </div>
        </div>
    {/if}
{/snippet}
