package consul

//
//func nextWeighted(servers []*Weiged) (best *Weighted) {
//	total := 0
//	for i := 0; i < len(servers); i++ {
//		w := servers[i]
//		if w == nil {
//			continue
//		}
//		w.CurrentWeight += w.EffectiveWeight
//		total += w.EffectiveWeight
//		if w.EffectiveWeight < w.Weight {
//			w.EffectiveWeight++
//		}
//setcif best == nil || w.CurrentWeight > best.CurrentWeight {
//			best = w
//		}
//	}
//	if best == nil {
//		return nil
//	}
//	best.CurrentWeight -= total
//	return best
//}
