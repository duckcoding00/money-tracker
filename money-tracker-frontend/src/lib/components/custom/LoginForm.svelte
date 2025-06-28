<script lang="ts">
	import Button from '../ui/button/button.svelte';
	import Input from '../ui/input/input.svelte';
	import Label from '../ui/label/label.svelte';

	type props = {
		email?: string;
		password?: string;
	};

	let { email, password }: props = $props();

	function cancelForm() {
		email = '';
		password = '';
	}
</script>

<div>
	<form action="" method="post" class="space-y-4">
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

		<div class="flex flex-row gap-3">
			<Button class="h-11 bg-gray-800 font-semibold text-white hover:bg-gray-900" type="submit"
				>Lets dive in</Button
			>
			<Button
				class="h-11 border-gray-400 font-medium "
				variant="outline"
				onclick={() => cancelForm()}>Start Over</Button
			>
		</div>
	</form>
</div>
