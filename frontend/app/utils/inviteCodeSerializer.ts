import { type InviteCode, type FormattedInviteCode } from "~/types/inviteCode"

export default (inviteCodes: InviteCode[]): FormattedInviteCode[] =>
{
	return inviteCodes.map(inviteCode =>
		{
			return {
				id        : inviteCode.id,
				code      : inviteCode.code,
				uses      : inviteCode.uses,
				maxUses   : inviteCode.max_uses,
				createdAt : inviteCode.created_at
			}
		}
	)
}