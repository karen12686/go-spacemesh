// Code generated by github.com/spacemeshos/go-scale/scalegen. DO NOT EDIT.

// nolint
package types

import (
	"github.com/spacemeshos/go-scale"
)

func (t *Ballot) EncodeScale(enc *scale.Encoder) (total int, err error) {
	{
		n, err := t.BallotMetadata.EncodeScale(enc)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := t.InnerBallot.EncodeScale(enc)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeByteArray(enc, t.Signature[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := t.Votes.EncodeScale(enc)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeStructSliceWithLimit(enc, t.EligibilityProofs, 500)
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func (t *Ballot) DecodeScale(dec *scale.Decoder) (total int, err error) {
	{
		n, err := t.BallotMetadata.DecodeScale(dec)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := t.InnerBallot.DecodeScale(dec)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.DecodeByteArray(dec, t.Signature[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := t.Votes.DecodeScale(dec)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		field, n, err := scale.DecodeStructSliceWithLimit[VotingEligibility](dec, 500)
		if err != nil {
			return total, err
		}
		total += n
		t.EligibilityProofs = field
	}
	return total, nil
}

func (t *BallotMetadata) EncodeScale(enc *scale.Encoder) (total int, err error) {
	{
		n, err := t.Layer.EncodeScale(enc)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeByteArray(enc, t.MsgHash[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func (t *BallotMetadata) DecodeScale(dec *scale.Decoder) (total int, err error) {
	{
		n, err := t.Layer.DecodeScale(dec)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.DecodeByteArray(dec, t.MsgHash[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func (t *InnerBallot) EncodeScale(enc *scale.Encoder) (total int, err error) {
	{
		n, err := scale.EncodeByteArray(enc, t.AtxID[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeByteArray(enc, t.OpinionHash[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeByteArray(enc, t.RefBallot[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeOption(enc, t.EpochData)
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func (t *InnerBallot) DecodeScale(dec *scale.Decoder) (total int, err error) {
	{
		n, err := scale.DecodeByteArray(dec, t.AtxID[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.DecodeByteArray(dec, t.OpinionHash[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.DecodeByteArray(dec, t.RefBallot[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		field, n, err := scale.DecodeOption[EpochData](dec)
		if err != nil {
			return total, err
		}
		total += n
		t.EpochData = field
	}
	return total, nil
}

func (t *Votes) EncodeScale(enc *scale.Encoder) (total int, err error) {
	{
		n, err := scale.EncodeByteArray(enc, t.Base[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeStructSliceWithLimit(enc, t.Support, 10000)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeStructSliceWithLimit(enc, t.Against, 10000)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeStructSliceWithLimit(enc, t.Abstain, 10000)
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func (t *Votes) DecodeScale(dec *scale.Decoder) (total int, err error) {
	{
		n, err := scale.DecodeByteArray(dec, t.Base[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		field, n, err := scale.DecodeStructSliceWithLimit[Vote](dec, 10000)
		if err != nil {
			return total, err
		}
		total += n
		t.Support = field
	}
	{
		field, n, err := scale.DecodeStructSliceWithLimit[Vote](dec, 10000)
		if err != nil {
			return total, err
		}
		total += n
		t.Against = field
	}
	{
		field, n, err := scale.DecodeStructSliceWithLimit[LayerID](dec, 10000)
		if err != nil {
			return total, err
		}
		total += n
		t.Abstain = field
	}
	return total, nil
}

func (t *Vote) EncodeScale(enc *scale.Encoder) (total int, err error) {
	{
		n, err := scale.EncodeByteArray(enc, t.ID[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := t.LayerID.EncodeScale(enc)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeCompact64(enc, uint64(t.Height))
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func (t *Vote) DecodeScale(dec *scale.Decoder) (total int, err error) {
	{
		n, err := scale.DecodeByteArray(dec, t.ID[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := t.LayerID.DecodeScale(dec)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		field, n, err := scale.DecodeCompact64(dec)
		if err != nil {
			return total, err
		}
		total += n
		t.Height = uint64(field)
	}
	return total, nil
}

func (t *Opinion) EncodeScale(enc *scale.Encoder) (total int, err error) {
	{
		n, err := scale.EncodeByteArray(enc, t.Hash[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := t.Votes.EncodeScale(enc)
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func (t *Opinion) DecodeScale(dec *scale.Decoder) (total int, err error) {
	{
		n, err := scale.DecodeByteArray(dec, t.Hash[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := t.Votes.DecodeScale(dec)
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func (t *EpochData) EncodeScale(enc *scale.Encoder) (total int, err error) {
	{
		n, err := scale.EncodeStructSliceWithLimit(enc, t.ActiveSet, 100000)
		if err != nil {
			return total, err
		}
		total += n
	}
	{
		n, err := scale.EncodeByteArray(enc, t.Beacon[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}

func (t *EpochData) DecodeScale(dec *scale.Decoder) (total int, err error) {
	{
		field, n, err := scale.DecodeStructSliceWithLimit[ATXID](dec, 100000)
		if err != nil {
			return total, err
		}
		total += n
		t.ActiveSet = field
	}
	{
		n, err := scale.DecodeByteArray(dec, t.Beacon[:])
		if err != nil {
			return total, err
		}
		total += n
	}
	return total, nil
}
