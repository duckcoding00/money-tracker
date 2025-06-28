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
		enhance?: SubmitFunction;
	};

	let {
		email,
		username,
		password,
		confirm_password,
		action,
		enhance: registerEnhance
	}: props = $props();

	function cancelForm() {
		email = '';
		username = '';
		password = '';
		confirm_password = '';
	}
</script>

<div>
	<form {action} method="post" class="space-y-4" use:enhance={registerEnhance}>
		<div class="space-y-2">
			<Label for="email" class="text-base font-semibold text-gray-800">Your Email Magic</Label>
			<Input
				id="email"
				type="email"
				placeholder="Enter your email"
				class="h-11 border-gray-300 focus:border-gray-500 focus:ring-gray-500"
				bind:value={email}
			/>
			<p class="text-sm font-medium text-gray-400">We'll send you amazing updates here!</p>
		</div>
		<div class="space-y-2">
			<Label for="username" class="text-base font-semibold text-gray-800"
				>Your Awesome Username</Label
			>
			<Input
				id="username"
				type="text"
				placeholder="Enter your Username"
				class="h-11 border-gray-300 focus:border-gray-500 focus:ring-gray-500"
				bind:value={username}
			/>
			<p class="text-sm font-medium text-gray-400">Make it uniquely YOU!</p>
		</div>
		<div class="space-y-2">
			<Label for="password" class="text-base font-semibold text-gray-800"
				>Super Secret Password</Label
			>
			<Input
				id="password"
				type="password"
				placeholder="Enter your Password"
				class="h-11 border-gray-300 focus:border-gray-500 focus:ring-gray-500"
				bind:value={password}
			/>
			<p class="text-sm font-medium text-gray-400">Make it strong like a superhero!</p>
		</div>
		<div class="space-y-2">
			<Label for="confirm_password" class="text-base font-semibold text-gray-800"
				>Confirm Your Super Password</Label
			>
			<Input
				id="confirm_password"
				type="password"
				placeholder="Enter your Confirm Password"
				class="h-11 border-gray-300 focus:border-gray-500 focus:ring-gray-500"
				bind:value={confirm_password}
			/>
			<p class="text-sm font-medium text-gray-400">Double-check for extra awesomeness!</p>
		</div>
		<div class="flex flex-row gap-3">
			<Button class="h-11 bg-gray-800 font-semibold text-white hover:bg-gray-900" type="submit"
				>Join the Adventure!</Button
			>
			<Button
				class="h-11 border-gray-400 font-medium "
				variant="outline"
				onclick={() => cancelForm()}>Start Over</Button
			>
		</div>
	</form>
</div>
