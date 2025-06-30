import { fail, isRedirect, redirect, type Actions } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load = (async () => {
	return {};
}) satisfies PageServerLoad;

export const actions: Actions = {
	verify: async ({ fetch, request }) => {
		const form = await request.formData();
		const token = form.get('otp') as string;
		const username = form.get('username') as string;

		const body = JSON.stringify({
			token,
			username
		});

		try {
			const response = await fetch('http://127.0.0.1:8080/api/v1/token/verify', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					Accept: 'application/json'
				},
				body: body
			});

			const result = await response.json();
			console.log(result);
			let otp = result.data;

			if (!response.ok) {
				if (result.details === 'redis value didnt exists') {
					return fail(400, { message: 'Are you sure that is the correct OTP?' });
				}
				return fail(response.status, result);
			}

			throw redirect(303, `/forgot-password/new?token=${otp}&username=${username}`);
		} catch (error) {
			if (isRedirect(error)) throw error;

			return fail(500, { message: 'Server error occurred' });
		}
	}
};
