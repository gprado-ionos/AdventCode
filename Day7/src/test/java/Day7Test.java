import java.util.List;
import java.util.Map;
import java.util.stream.Stream;
import org.assertj.core.api.Assertions;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

public class Day7Test {

  public static Stream<Arguments> cardHandAndTypes() {
    return Stream.of(
        Arguments.of("32T3K", CardType.ONE_PAIR),
        Arguments.of("T55J5", CardType.THREE_OF_A_KIND),
        Arguments.of("KK677", CardType.TWO_PAIR),
        Arguments.of("QQQJA", CardType.THREE_OF_A_KIND),
        Arguments.of("KTJJT", CardType.TWO_PAIR),
        Arguments.of("23332", CardType.FULL_HOUSE),
        Arguments.of("44644", CardType.FOUR_OF_A_KIND),
        Arguments.of("AAKAA", CardType.FOUR_OF_A_KIND),
        Arguments.of("34567", CardType.HIGH_CARD),
        Arguments.of("65432", CardType.HIGH_CARD),
        Arguments.of("AAAAA", CardType.FIVE_OF_A_KIND)
    );
  }

  @ParameterizedTest
  @MethodSource("cardHandAndTypes")
  void test__givenCardHand__thenReturnCardType(String cardHand, CardType cardType) {
    CardType ct = CardHandBuilder.builder().withHand(cardHand).build().cardType();
    Assertions.assertThat(cardType).isEqualTo(ct);
  }

  public static Stream<Arguments> cardHandAndTypesJoker() {
    return Stream.of(
        Arguments.of("32T3K", CardType.ONE_PAIR),
        Arguments.of("T55J5", CardType.FOUR_OF_A_KIND),
        Arguments.of("KK677", CardType.TWO_PAIR),
        Arguments.of("QQQJA", CardType.FOUR_OF_A_KIND),
        Arguments.of("KTJJT", CardType.FOUR_OF_A_KIND),
        Arguments.of("23332", CardType.FULL_HOUSE),
        Arguments.of("44J44", CardType.FIVE_OF_A_KIND),
        Arguments.of("JJKAA", CardType.FOUR_OF_A_KIND),
        Arguments.of("3J567", CardType.ONE_PAIR),
        Arguments.of("65432", CardType.HIGH_CARD),
        Arguments.of("AAAAA", CardType.FIVE_OF_A_KIND)
    );
  }

  @ParameterizedTest
  @MethodSource("cardHandAndTypesJoker")
  void test__givenCardHand__thenReturnCardTypeJoker(String cardHand, CardType cardType) {
    CardType ct = CardHandBuilder.builder().withHand(cardHand).build().cardTypeJoker();
    Assertions.assertThat(cardType).isEqualTo(ct);
  }

  @Test
  void test__givenCardHands__returnPoints_Part1() {
    List<String> cardHands = List.of(
        "32T3K 765",
        "T55J5 684",
        "KK677 28",
        "KTJJT 220",
        "QQQJA 483"
    );
    Day7 day7 = new Day7();
    List<CardHand> cardsAndPoints = day7.parseCardHands(cardHands);
    Map<CardType, List<CardHand>> cardsGroupedByType = day7.groupByCardType(cardsAndPoints, false);
    Long points = day7.getPoints(cardsGroupedByType);
    Assertions.assertThat(points).isEqualTo(6440);
  }

  @Test
  void test__givenCardHands__returnPoints_Part1_ex2() {
    List<String> cardHands = List.of(
        "65432 5",
        "AKJT9 4",
        "QQJ7J 8",
        "A76AA 10",
        "88A88 4",
        "88888 4"
    );
    Day7 day7 = new Day7();
    List<CardHand> cardsAndPoints = day7.parseCardHands(cardHands);
    Map<CardType, List<CardHand>> cardsGroupedByType = day7.groupByCardType(cardsAndPoints, false);
    Long points = day7.getPoints(cardsGroupedByType);
    Assertions.assertThat(points).isEqualTo(121);
  }

  @Test
  void test__givenCardHandString__returnCardHand() {
    String cardHandString = "JAAA7 340";
    Day7 day7 = new Day7();
    CardHand cardHand = day7.getCardHand(cardHandString);
    Assertions.assertThat(cardHand.cardHand()).isEqualTo("JAAA7");
    Assertions.assertThat(cardHand.points()).isEqualTo(340);
  }

  @Test
  void test__givenCardHands__returnPoints_Part2() {
    List<String> cardHands = List.of(
        "32T3K 765",
        "T55J5 684",
        "KK677 28",
        "KTJJT 220",
        "QQQJA 483"
    );
    Day7 day7 = new Day7();
    List<CardHand> cardsAndPoints = day7.parseCardHands(cardHands);
    Map<CardType, List<CardHand>> cardsGroupedByType = day7.groupByCardType(cardsAndPoints, true);
    Long points = day7.getPoints(cardsGroupedByType);
    Assertions.assertThat(points).isEqualTo(5905);
  }
}
