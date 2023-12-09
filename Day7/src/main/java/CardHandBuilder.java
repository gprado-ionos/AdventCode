import java.util.Arrays;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Optional;
import java.util.stream.Collectors;

public class CardHandBuilder {

  private String hand;
  private String handJokerApplied;
  private Integer points;
  private CardType cardTypeJoker;
  private CardType cardType;

  private CardHandBuilder() {
  }

  public static CardHandBuilder builder() {
    return new CardHandBuilder();
  }

  public CardHandBuilder withHand(String hand) {
    this.hand = hand;
    return this;
  }

  public CardHandBuilder withPoints(Integer integer) {
    this.points = integer;
    return this;
  }

  public CardHand build() {

    cardType = findCardType();
    applyJoker();
    return new CardHand(hand, points, cardType, cardTypeJoker);
  }

  private void applyJoker() {
    Map<String, Long> collectChars = Arrays.stream(
        this.hand.toLowerCase().split("")).collect(
        Collectors.groupingBy(c -> c, Collectors.counting()));

    Optional<Entry<String, Long>> hasJoker = collectChars.entrySet().stream()
        .filter(entry -> entry.getKey().equalsIgnoreCase("j")).findFirst();

    if (hasJoker.isPresent()) {
      if (this.cardType == CardType.FIVE_OF_A_KIND) {
        this.cardTypeJoker = CardType.FIVE_OF_A_KIND;
        //this.handJokerApplied = this.hand;

      } else if (this.cardType == CardType.FOUR_OF_A_KIND) {
        this.cardTypeJoker = CardType.FIVE_OF_A_KIND;
        //this.handJokerApplied = this.hand.replace(hasJoker.get().getKey(), "A");
      } else if (this.cardType == CardType.FULL_HOUSE) {
        this.cardTypeJoker = CardType.FIVE_OF_A_KIND;
        //this.handJokerApplied = this.hand.replace(hasJoker.get().getKey(), "A");
      } else if (this.cardType == CardType.THREE_OF_A_KIND) {
        this.cardTypeJoker = CardType.FOUR_OF_A_KIND;
        //this.handJokerApplied = this.hand.replace(hasJoker.get().getKey(), "A");
      } else if (this.cardType == CardType.TWO_PAIR) {
        if (hasJoker.get().getValue() == 1) {
          this.cardTypeJoker = CardType.FULL_HOUSE;
          //collectChars.entrySet().stream().filter(entry -> entry.getKey())
        } else {
          this.cardTypeJoker = CardType.FOUR_OF_A_KIND;
        }
      } else if (this.cardType == CardType.ONE_PAIR) {
        this.cardTypeJoker = CardType.THREE_OF_A_KIND;

      } else if (this.cardType == CardType.HIGH_CARD) {
        this.cardTypeJoker = CardType.ONE_PAIR;
      }
    } else {
      this.cardTypeJoker = this.cardType;
    }
  }

  private CardType findCardType() {
    Map<String, Long> collectChars = Arrays.stream(
        this.hand.toLowerCase().split("")).collect(
        Collectors.groupingBy(c -> c, Collectors.counting()));

    CardType type = null;
    if (collectChars.size() == 1) {
      type = CardType.FIVE_OF_A_KIND;
    } else if (collectChars.size() == 2) {
      for (Map.Entry<String, Long> entry : collectChars.entrySet()) {
        if (entry.getValue() == 1) {
          continue;
        } else if (entry.getValue() == 4) {
          type = CardType.FOUR_OF_A_KIND;
        } else {
          type = CardType.FULL_HOUSE;
        }
      }
    } else if (collectChars.size() == 3) {
      for (Map.Entry<String, Long> entry : collectChars.entrySet()) {
        if (entry.getValue() == 1) {
          continue;
        } else if (entry.getValue() == 3) {
          type = CardType.THREE_OF_A_KIND;
        } else {
          type = CardType.TWO_PAIR;
        }
      }
    } else if (collectChars.size() == 4) {
      type = CardType.ONE_PAIR;
    } else {
      type = CardType.HIGH_CARD;
    }

    return type;
  }
}
