import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Comparator;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

public class Day7 {

  public static final String VALUE_ORDER = "AKQJT98765432";

  public static final String VALUE_ORDER_PT2 = "AKQT98765432J";
  public static void main(String[] args) throws IOException {
    Day7 day7 = new Day7();

    List<String> allLines = Files.readAllLines(Paths.get(
        "/home/gprado/Development/private_repo/AdventCode/Day7/src/main/resources/input.txt"));

    List<CardHand> cardHands = day7.parseCardHands(allLines);
    Map<CardType, List<CardHand>> cardsGroupedByType = day7.groupByCardType(cardHands, true);
    Long points = day7.getPoints(cardsGroupedByType);
    System.out.println(points);

  }

  public CardHand getCardHand(String cardHandString) {
    String[] cardHandArray = cardHandString.split(" ");
    CardHand cardHand = CardHandBuilder.builder().withHand(cardHandArray[0]).withPoints(Integer.valueOf(cardHandArray[1])).build();
    return cardHand;
  }

  public List<CardHand> parseCardHands(List<String> cardHands) {
    return cardHands.stream().map(this::getCardHand).toList();
  }

  public Map<CardType, List<CardHand>> groupByCardType(List<CardHand> cardsAndPoints, boolean useJoker) {
    if (useJoker) {
      return cardsAndPoints.stream().collect(Collectors.groupingBy(CardHand::cardTypeJoker));
    }
    return cardsAndPoints.stream().collect(Collectors.groupingBy(CardHand::cardType));
  }

  public Long getPoints(Map<CardType, List<CardHand>> cardsGroupedByType) {
    Integer points = 1;
    Long totalPoints = 0l;
    Comparator<CardHand> reverseCardHandComparator =
        (ct1, ct2) -> {
          int pos1 = 0;
          int pos2 = 0;
          for (int i = 0; i < Math.min(ct1.cardHand().length(), ct2.cardHand().length()) && pos1 == pos2; i++) {
            pos1 = VALUE_ORDER_PT2.indexOf(ct1.cardHand().charAt(i));
            pos2 = VALUE_ORDER_PT2.indexOf(ct2.cardHand().charAt(i));
          }
          return pos2 - pos1;
        };
    for (CardType cardType : CardType.values()) {
      if (cardsGroupedByType.containsKey(cardType)) {
        List<CardHand> collected = cardsGroupedByType.get(cardType).stream()
            .sorted(reverseCardHandComparator).collect(Collectors.toList());
        for (CardHand cardHand : collected) {
          totalPoints += cardHand.points() * points;
          points++;
        }

      }
    }
    return totalPoints;

  }
}
